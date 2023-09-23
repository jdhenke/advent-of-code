package day4

import (
	"bytes"
	"io"
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
	return day4(r, matchesPart1)
}

/*
Part2 Prompt

--- Part Two ---
An Elf just remembered one more important detail: the two adjacent matching
digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the
following are now true:

- 112233 meets these criteria because the digits never decrease and all
repeated digits are exactly two digits long.
- 123444 no longer meets the criteria (the repeated 44 is part of a larger
group of 444).
- 111122 meets the criteria (even though 1 is repeated more than twice, it
still contains a double 22).

How many different passwords within the range given in your puzzle input meet
all of the criteria?
*/
func Part2(r io.Reader) (answer int, err error) {
	return day4(r, matchesPart2)
}

func day4(r io.Reader, matchesFunc func(digits []int) bool) (answer int, err error) {
	b, err := io.ReadAll(r)
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
		if matchesFunc(toDigits(x)) {
			answer++
		}
	}
	return answer, nil
}

func matchesPart1(digits []int) bool {
	double := false
	for i, d := range digits {
		if i-1 >= 0 {
			if digits[i-1] > d {
				return false
			}
			if digits[i-1] == d {
				double = true
			}
		}
	}
	return double
}

func matchesPart2(digits []int) bool {
	strictDouble := false
	for i, d := range digits {
		if i-1 >= 0 {
			if digits[i-1] > d {
				return false
			}
			if digits[i-1] == d {
				if (i-2 < 0 || digits[i-2] != d) &&
					((i+1) >= len(digits) || digits[i+1] != d) {
					strictDouble = true
				}
			}
		}
	}
	return strictDouble
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
