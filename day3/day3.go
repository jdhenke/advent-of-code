package day3

import (
	"advent-of-code/input"
	"fmt"
	"io"
)

func Part1(r io.Reader) (ans string, err error) {
	// aggregate frequencies
	var freqs []map[string]int
	err = input.ForEachLine(r, func(line string) error {
		for i := 0; i < len(line); i++ {
			if i >= len(freqs) {
				freqs = append(freqs, map[string]int{})
			}
			c := line[len(line)-i-1 : len(line)-i]
			freqs[i][c]++
		}
		return nil
	})
	// create gamma / epsilon things
	var gamma, epsilon int
	for i := range freqs {
		fs := freqs[len(freqs)-i-1]
		gamma <<= 1
		epsilon <<= 1
		if fs["1"] > fs["0"] {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	return fmt.Sprint(gamma * epsilon), nil
}

func Part2(r io.Reader) (ans string, err error) {
	panic("todo")
}
