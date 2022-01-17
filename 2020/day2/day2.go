package day2

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 2: Password Philosophy ---
Your flight departs in a few days from the coastal airport; the easiest way
down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day.
"Something's wrong with our computers; we can't log in!" You ask if you can
take a look.

Their password database seems to be a little corrupted: some of the passwords
wouldn't have been allowed by the Official Toboggan Corporate Policy that was
in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of
passwords (according to the corrupted database) and the corporate policy when
that password was set.

For example, suppose you have the following list:

    1-3 a: abcde
    1-3 b: cdefg
    2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy
indicates the lowest and highest number of times a given letter must appear for
the password to be valid. For example, 1-3 a means that the password must
contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is
not; it contains no instances of b, but needs at least 1. The first and third
passwords are valid: they contain one a or nine c, both within the limits of
their respective policies.

How many passwords are valid according to their policies?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day2(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day2(r)
}

// 1-5 d: ddddbd
var re = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func day2(r io.Reader) (answer int, err error) {
	if err := input.ForEachLine(r, func(line string) error {
		parts := re.FindStringSubmatch(line)
		if parts == nil {
			return fmt.Errorf("bad line: %v", line)
		}
		min, max := num(parts[1]), num(parts[2])
		c := parts[3]
		p := parts[4]
		n := 0
		for i := range p {
			if p[i:i+1] == c {
				n++
			}
		}
		if n >= min && n <= max {
			answer++
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return answer, nil
}

func num(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
