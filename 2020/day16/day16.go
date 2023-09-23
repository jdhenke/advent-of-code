package day16

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 16: Ticket Translation ---
As you're walking to yet another connecting flight, you realize that one of the
legs of your re-routed trip coming up is on a high-speed train. However, the
train ticket you were given is in a language you don't understand. You should
probably figure out what it says before you get to the train station after the
next flight.

Unfortunately, you can't actually read the words on the ticket. You can,
however, read the numbers, and so you figure out the fields these tickets must
have and the valid ranges for values in those fields.

You collect the rules for ticket fields, the numbers on your ticket, and the
numbers on other nearby tickets for the same train service (via the airport
security cameras) together into a single document you can reference (your
puzzle input).

The rules for ticket fields specify a list of fields that exist somewhere on
the ticket and the valid ranges of values for each field. For example, a rule
like class: 1-3 or 5-7 means that one of the fields in every ticket is named
class and can be any value in the ranges 1-3 or 5-7 (inclusive, such that 3 and
5 are both valid in this field, but 4 is not).

Each ticket is represented by a single line of comma-separated values. The
values are the numbers on the ticket in the order they appear; every ticket has
the same format. For example, consider this ticket:

	.--------------------------------------------------------.
	| ????: 101    ?????: 102   ??????????: 103     ???: 104 |
	|                                                        |
	| ??: 301  ??: 302             ???????: 303      ??????? |
	| ??: 401  ??: 402           ???? ????: 403    ????????? |
	'--------------------------------------------------------'

Here, ? represents text in a language you don't understand. This ticket might
be represented as 101,102,103,104,301,302,303,401,402,403; of course, the
actual train tickets you're looking at are much more complicated. In any case,
you've extracted just the numbers in such a way that the first number is always
the same specific field, the second number is always a different specific
field, and so on - you just don't know what each position actually means!

Start by determining which tickets are completely invalid; these are tickets
that contain values which aren't valid for any field. Ignore your ticket for
now.

For example, suppose you have the following notes:

	class: 1-3 or 5-7
	row: 6-11 or 33-44
	seat: 13-40 or 45-50

	your ticket:
	7,1,14

	nearby tickets:
	7,3,47
	40,4,50
	55,2,20
	38,6,12

It doesn't matter which position corresponds to which field; you can identify
invalid nearby tickets by considering only whether tickets contain values that
are not valid for any field. In this example, the values on the first nearby
ticket are all valid for at least one field. This is not true of the other
three nearby tickets: the values 4, 55, and 12 are are not valid for any field.
Adding together all of the invalid values produces your ticket scanning error
rate: 4 + 55 + 12 = 71.

Consider the validity of the nearby tickets you scanned. What is your ticket
scanning error rate?
*/
func Part1(r io.Reader) (answer int, err error) {
	invalid, _, err := day16(r, false)
	return invalid, err
}

/*
Part2 Prompt

--- Part Two ---
Now that you've identified which tickets contain invalid values, discard those
tickets entirely. Use the remaining valid tickets to determine which field is
which.

Using the valid ranges for each field, determine what order the fields appear
on the tickets. The order is consistent between all tickets: if seat is the
third field, it is the third field on every ticket, including your ticket.

For example, suppose you have the following notes:

	class: 0-1 or 4-19
	row: 0-5 or 8-19
	seat: 0-13 or 16-19

	your ticket:
	11,12,13

	nearby tickets:
	3,9,18
	15,1,5
	5,14,9

Based on the nearby tickets in the above example, the first position must be
row, the second position must be class, and the third position must be seat;
you can conclude that in your ticket, class is 12, row is 11, and seat is 13.

Once you work out which field is which, look for the six fields on your ticket
that start with the word departure. What do you get if you multiply those six
values together?
*/
func Part2(r io.Reader) (answer int, err error) {
	_, ticket, err := day16(r, true)
	if err != nil {
		return 0, err
	}
	answer = 1
	for s, v := range ticket {
		if strings.HasPrefix(s, "departure") {
			answer *= v
		}
	}
	return answer, nil
}

// class: 1-3 or 5-7
var re = regexp.MustCompile(`(.+): (\d+)-(\d+) or (\d+)-(\d+)`)

type ruleSet map[string]rule

func (rs ruleSet) Matches(x int) bool {
	for _, r := range rs {
		if r.Matches(x) {
			return true
		}
	}
	return false
}

type rule [2][2]int

func (r rule) Matches(x int) bool {
	return (x >= r[0][0] && x <= r[0][1]) || (x >= r[1][0] && x <= r[1][1])
}

func day16(r io.Reader, doSolve bool) (invalid int, ticket map[string]int, err error) {
	var mode int
	rules := make(ruleSet)
	var valid [][]int
	var myTicket []int
	if err := input.ForEachLine(r, func(line string) error {
		if line == "" {
			mode++
			return nil
		}
		switch mode {
		case 0:
			parts := re.FindStringSubmatch(line)
			if parts == nil {
				return fmt.Errorf("bad rule line: %v", line)
			}
			nums := mustNums(parts[2:]...)
			rules[parts[1]] = rule{{nums[0], nums[1]}, {nums[2], nums[3]}}
		case 1:
			if line == "your ticket:" {
				return nil
			}
			myTicket = mustNums(strings.Split(line, ",")...)
		case 2:
			if line == "nearby tickets:" {
				return nil
			}
			parts := strings.Split(line, ",")
			nums := mustNums(parts...)
			isValid := true
			for _, num := range nums {
				if !rules.Matches(num) {
					invalid += num
					isValid = false
				}
			}
			if isValid {
				valid = append(valid, nums)
			}
		}
		return nil
	}); err != nil {
		return 0, nil, err
	}
	if doSolve {
		ticket, err = solve(rules, myTicket, valid)
		if err != nil {
			return 0, nil, err
		}
	}
	return invalid, ticket, nil
}

func solve(rules ruleSet, myTicket []int, tickets [][]int) (map[string]int, error) {
	// start with all possible options
	options := make(map[int]map[string]bool)
	for i := range myTicket {
		opts := make(map[string]bool)
		for s := range rules {
			opts[s] = true
		}
		options[i] = opts
	}
	done := make(map[int]bool)
	solved := func() bool {
		for _, opts := range options {
			if len(opts) != 1 {
				return false
			}
		}
		return true
	}
	matches := func(name string, i int) bool {
		for _, ticket := range tickets {
			if !rules[name].Matches(ticket[i]) {
				return false
			}
		}
		return true
	}
	// remove options until it's solved
	for !solved() {
		changed := false
		// see if any columns contain a valid that would not match one of its existing options and if so, remove that
		// option
		for i := 0; i < len(myTicket); i++ {
			if len(options[i]) == 1 {
				continue
			}
			for s := range options[i] {
				if !matches(s, i) {
					delete(options[i], s)
					changed = true
				}
			}
		}
		// see if any new columns have been solved and if so, remove that column's answer from all other columns'
		// options
		for i, opts := range options {
			if done[i] {
				continue
			}
			if len(opts) == 1 {
				ans := only(opts)
				done[i] = true
				for j, otherOpts := range options {
					if i == j {
						continue
					}
					delete(otherOpts, ans)
				}
				changed = true
			}
		}
		if !changed {
			return nil, fmt.Errorf("no solution")
		}
	}
	// convert myTicket into text answers
	out := make(map[string]int)
	for i := range myTicket {
		out[only(options[i])] = myTicket[i]
	}
	return out, nil
}

func mustNums(all ...string) []int {
	var nums []int
	for _, s := range all {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, x)
	}
	return nums
}

func only(m map[string]bool) string {
	for s := range m {
		return s
	}
	panic("empty map")
}
