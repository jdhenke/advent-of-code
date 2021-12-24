package day1

import (
	"advent-of-code/input"
	"fmt"
	"io"
)

const size = 3

func Part1(r io.Reader) (answer string, err error) {
	var (
		i         = 0
		recent    = 0
		increases = 0
	)
	if err := input.ScanInt(r, func(x int) {
		i++
		if i > 1 && x > recent {
			increases++
		}
		recent = x
	}); err != nil {
		return "", err
	}
	return fmt.Sprint(increases), nil
}

func Part2(r io.Reader) (answer string, err error) {
	var (
		i         = 0
		recent    [size]int
		increases = 0
	)
	if err := input.ScanInt(r, func(x int) {
		i++
		if i > 3 && x > recent[i%3] {
			increases++
		}
		recent[i%3] = x
	}); err != nil {
		return "", err
	}
	return fmt.Sprint(increases), nil
}
