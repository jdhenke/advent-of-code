package day1

import (
	"advent-of-code/input"
	"io"
)

const size = 3

func Part1(r io.Reader) (answer int, err error) {
	var (
		i         = 0
		recent    = 0
		increases = 0
	)
	if err := input.ForEachInt(r, func(x int) {
		i++
		if i > 1 && x > recent {
			increases++
		}
		recent = x
	}); err != nil {
		return 0, err
	}
	return increases, nil
}

func Part2(r io.Reader) (answer int, err error) {
	var (
		i         = 0
		recent    [size]int
		increases = 0
	)
	if err := input.ForEachInt(r, func(x int) {
		i++
		if i > 3 && x > recent[i%3] {
			increases++
		}
		recent[i%3] = x
	}); err != nil {
		return 0, err
	}
	return increases, nil
}
