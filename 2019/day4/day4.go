package day4

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
)

/*
Part1 Prompt

--- Day 4: Secure Container ---
You arrive at the Venus fuel depot only to discover it's protected by a
password. The Elves had written the password on a sticky note, but someone
threw it out.

However, they do remember a few key facts about the password:

- It is a six-digit number.
- The value is within the range given in your puzzle input.
- Two adjacent digits are the same (like 22 in 122345).
- Going from left to right, the digits never decrease; they only ever increase
or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

- 111111 meets these criteria (double 11, never decreases).
- 223450 does not meet these criteria (decreasing pair of digits 50).
- 123789 does not meet these criteria (no double).

How many different passwords within the range given in your puzzle input meet
these criteria?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day4(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day4(r)
}

func day4(r io.Reader) (answer int, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	parts := bytes.Split(bytes.TrimSpace(b), []byte{'-'})
	beg, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return 0, err
	}
	end, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		return 0, err
	}
	for x := beg; x <= end; x++ {
		if matches(x) {
			answer++
		}
	}
	return answer, nil
}

func matches(x int) bool {
	digits := toDigits(x)
	double := false
	for i, d := range digits {
		if i-1 >= 0 {
			if digits[i-1] == d {
				double = true
			}
			if digits[i-1] > d {
				return false
			}
		}
	}
	return double
}

func toDigits(x int) []int {
	var reverse []int
	for x > 0 {
		reverse = append(reverse, x%10)
		x /= 10
	}
	var out []int
	for i := 0; i < len(reverse); i++ {
		out = append(out, reverse[len(reverse)-i-1])
	}
	return out
}
