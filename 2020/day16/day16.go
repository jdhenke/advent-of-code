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
	return day16(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day16(r)
}

// class: 1-3 or 5-7

var re = regexp.MustCompile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`)

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

func day16(r io.Reader) (answer int, err error) {
	var mode int
	rules := make(ruleSet)
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
			// IGNORE FOR NOW
		case 2:
			if line == "nearby tickets:" {
				return nil
			}
			parts := strings.Split(line, ",")
			for _, num := range mustNums(parts...) {
				if !rules.Matches(num) {
					answer += num
				}
			}
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return answer, nil
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
