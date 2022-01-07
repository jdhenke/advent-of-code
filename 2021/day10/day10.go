package day10

import (
	"advent-of-code/input"
	"io"
	"sort"
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

var part2Points = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func Part2(r io.Reader) (ans int, err error) {
	var scores []int
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
					return nil
				}
			}
		}
		score := 0
		for i := range stack {
			c := stack[len(stack)-i-1]
			score *= 5
			score += part2Points[c]
		}
		scores = append(scores, score)
		return nil
	}); err != nil {
		return 0, err
	}
	sort.Ints(scores)
	return scores[len(scores)/2], nil
}
