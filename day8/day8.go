package day8

import (
	"advent-of-code/input"
	"io"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	num1478 := 0
	if err := input.ForEachLine(r, func(line string) error {
		outputStrings := strings.SplitN(line, " | ", 2)[1]
		for _, s := range strings.Split(outputStrings, " ") {
			if map[int]bool{
				2: true,
				3: true,
				4: true,
				7: true,
			}[len(s)] {
				num1478++
			}
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return num1478, nil
}
