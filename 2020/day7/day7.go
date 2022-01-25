package day7

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 7: Handy Haversacks ---
You land at the regional airport in time for your next flight. In fact, it
looks like you'll even have time to grab some food: all flights are currently
delayed due to issues in luggage processing.

Due to recent aviation regulations, many rules (your puzzle input) are being
enforced about bags and their contents; bags must be color-coded and must
contain specific quantities of other color-coded bags. Apparently, nobody
responsible for these regulations considered how long they would take to
enforce!

For example, consider the following rules:

    light red bags contain 1 bright white bag, 2 muted yellow bags.
    dark orange bags contain 3 bright white bags, 4 muted yellow bags.
    bright white bags contain 1 shiny gold bag.
    muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
    shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
    dark olive bags contain 3 faded blue bags, 4 dotted black bags.
    vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
    faded blue bags contain no other bags.
    dotted black bags contain no other bags.

These rules specify the required contents for 9 bag types. In this example,
every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded
blue and 6 dotted black), and so on.

You have a shiny gold bag. If you wanted to carry it in at least one other bag,
how many different bag colors would be valid for the outermost bag? (In other
words: how many colors can, eventually, contain at least one shiny gold bag?)

In the above rules, the following options would be available to you:

- A bright white bag, which can hold your shiny gold bag directly.
- A muted yellow bag, which can hold your shiny gold bag directly, plus some
other bags.
- A dark orange bag, which can hold bright white and muted yellow bags, either
of which could then hold your shiny gold bag.
- A light red bag, which can hold bright white and muted yellow bags, either of
which could then hold your shiny gold bag.

So, in this example, the number of bag colors that can eventually contain at
least one shiny gold bag is 4.

How many bag colors can eventually contain at least one shiny gold bag? (The
list of rules is quite long; make sure you get all of it.)
*/
func Part1(r io.Reader) (answer int, err error) {
	rules, err := getRules(r)
	if err != nil {
		return 0, err
	}
	hasGold := memoHasGold()
	for k := range rules {
		if k != "shiny gold" && hasGold(rules, k) {
			answer++
		}
	}
	return answer, nil
}

/*
Part2 Prompt

--- Part Two ---
It's getting pretty expensive to fly these days - not because of ticket prices,
but because of the ridiculous number of bags you need to buy!

Consider again your shiny gold bag and the rules from the above example:

- faded blue bags contain 0 other bags.
- dotted black bags contain 0 other bags.
- vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black
bags.
- dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black
bags.

So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags
within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1 +
1*7 + 2 + 2*11 = 32 bags!

Of course, the actual rules have a small chance of going several levels deeper
than this example; be sure to count all of the bags, even if the nesting
becomes topologically impractical!

Here's another example:

    shiny gold bags contain 2 dark red bags.
    dark red bags contain 2 dark orange bags.
    dark orange bags contain 2 dark yellow bags.
    dark yellow bags contain 2 dark green bags.
    dark green bags contain 2 dark blue bags.
    dark blue bags contain 2 dark violet bags.
    dark violet bags contain no other bags.

In this example, a single shiny gold bag must contain 126 other bags.

How many individual bags are required inside your single shiny gold bag?
*/
func Part2(r io.Reader) (answer int, err error) {
	rules, err := getRules(r)
	if err != nil {
		return 0, err
	}
	numBags := memoNumBags()
	return numBags(rules, "shiny gold") - 1, nil
}

func memoHasGold() func(rules map[string]map[string]int, bag string) (ans bool) {
	memo := make(map[string]bool)
	var f func(rules map[string]map[string]int, bag string) (ans bool)
	f = func(rules map[string]map[string]int, bag string) (ans bool) {
		if ans, ok := memo[bag]; ok {
			return ans
		}
		defer func() { memo[bag] = ans }()
		for sub := range rules[bag] {
			if sub == "shiny gold" {
				return true
			}
			if f(rules, sub) {
				return true
			}
		}
		return false
	}
	return f
}

func memoNumBags() func(rules map[string]map[string]int, bag string) (ans int) {
	memo := make(map[string]int)
	var f func(rules map[string]map[string]int, bag string) (ans int)
	f = func(rules map[string]map[string]int, bag string) (ans int) {
		if ans, ok := memo[bag]; ok {
			return ans
		}
		defer func() { memo[bag] = ans }()
		ans++
		for sub, count := range rules[bag] {
			ans += count * f(rules, sub)
		}
		return ans
	}
	return f
}

func getRules(r io.Reader) (map[string]map[string]int, error) {
	rules := make(map[string]map[string]int)
	if err := input.ForEachLine(r, func(line string) error {
		parts := strings.Split(line, " bags contain ")
		rules[parts[0]] = parseBags(parts[1])
		return nil
	}); err != nil {
		return nil, err
	}
	return rules, nil
}

// Options:
//
// 3 vibrant green bags, 4 plaid blue bags, 2 drab brown bags.
// 1 pale magenta bag.
// no other bags.
var re = regexp.MustCompile(`^(\d+) (\w+ \w+) bag(s?)(, )?(\.)?(.+)$`)

func parseBags(s string) map[string]int {
	bags := make(map[string]int)
	for m := re.FindStringSubmatch(s); m != nil; m = re.FindStringSubmatch(m[6]) {
		x, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		bags[m[2]] = x
	}
	return bags
}
