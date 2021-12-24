package day3

import (
	"advent-of-code/input"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	// aggregate frequencies
	var freqs []map[string]int
	if err = input.ForEachLine(r, func(line string) error {
		for i := 0; i < len(line); i++ {
			if i >= len(freqs) {
				freqs = append(freqs, map[string]int{})
			}
			c := line[len(line)-i-1 : len(line)-i]
			freqs[i][c]++
		}
		return nil
	}); err != nil {
		return 0, err
	}
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
	return gamma * epsilon, nil
}

func Part2(r io.Reader) (ans int, err error) {
	rs, ok := r.(io.ReadSeeker)
	if !ok {
		return 0, fmt.Errorf("this solution requires a ReadSeeker")
	}
	o2Str, err := filter(rs, func(zeroes, ones int) string {
		if ones >= zeroes {
			return "1"
		}
		return "0"
	})
	if err != nil {
		return 0, err
	}
	co2Str, err := filter(rs, func(zeroes, ones int) string {
		if ones >= zeroes {
			return "0"
		}
		return "1"
	})
	if err != nil {
		return 0, err
	}
	o2, err := strconv.ParseInt(o2Str, 2, 64)
	if err != nil {
		return 0, err
	}
	co2, err := strconv.ParseInt(co2Str, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(o2 * co2), nil
}

// constantly filters until the last line that matches is left
func filter(rs io.ReadSeeker, suffix func(zeroes, ones int) string) (string, error) {
	prefix := ""
	for {
		if _, err := rs.Seek(0, io.SeekStart); err != nil {
			return "", fmt.Errorf("failed to seek to start: %v", err)
		}
		var numHits int
		var lastHit string
		var ones, zeroes int
		if err := input.ForEachLine(rs, func(line string) error {
			if strings.HasPrefix(line, prefix) {
				if len(prefix) < len(line) {
					if line[len(prefix):len(prefix)+1] == "1" {
						ones++
					} else {
						zeroes++
					}
				}
				numHits++
				lastHit = line
			}
			return nil
		}); err != nil {
			return "", err
		}
		if numHits == 0 {
			return "", fmt.Errorf("got to zero hits somehow")
		}
		if numHits == 1 {
			return lastHit, nil
		}
		prefix += suffix(zeroes, ones)
	}
}
