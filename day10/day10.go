package day10

import (
	"advent-of-code/input"
	"io"
)

var points = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var opening = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func Part1(r io.Reader) (ans int, err error) {
	total := 0
	if err := input.ForEachLine(r, func(line string) error {
		var stack []string
		for i := 0; i < len(line); i++ {
			c := line[i : i+1]
			switch c {
			case "(", "[", "{", "<":
				stack = append(stack, c)
			case ")", "]", "}", ">":
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if last != opening[c] {
					total += points[c]
					return nil
				}
			}
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return total, nil
}

func Part2(r io.Reader) (ans int, err error) {
	panic("tbd")
}
