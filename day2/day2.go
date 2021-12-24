package day2

import (
	"advent-of-code/input"
	"io"
)

func Part1(r io.Reader) (ans int, err error) {
	var (
		horizontalPos = 0
		depth         = 0
	)
	if err := input.ForEachCommand(r, func(cmd string, val int) {
		switch cmd {
		case "forward":
			horizontalPos += val
		case "up":
			depth -= val
		case "down":
			depth += val
		}
	}); err != nil {
		return 0, err
	}
	return horizontalPos * depth, nil
}

func Part2(r io.Reader) (ans int, err error) {
	var (
		aim           = 0
		horizontalPos = 0
		depth         = 0
	)
	if err := input.ForEachCommand(r, func(cmd string, val int) {
		switch cmd {
		case "forward":
			horizontalPos += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}); err != nil {
		return 0, err
	}
	return horizontalPos * depth, nil
}
