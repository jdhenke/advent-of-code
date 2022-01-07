package day8

import (
	"advent-of-code/input"
	"io"
	"sort"
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

func Part2(r io.Reader) (ans int, err error) {
	sum := 0
	if err := input.ForEachLine(r, func(line string) error {
		parts := strings.SplitN(line, " | ", 2)
		inputs, outputs := parts[0], parts[1]
		sum += solve(sortItems(strings.Split(inputs, " ")), sortItems(strings.Split(outputs, " ")))
		return nil
	}); err != nil {
		return 0, err
	}
	return sum, nil
}

func solve(inputs, outputs []string) int {
	options := make(map[string]map[int]bool)
	answers := make(map[int]string)
	for _, c := range inputs {
		switch len(c) {
		case 2:
			answers[1] = c
			options[c] = map[int]bool{1: true}
		case 3:
			answers[7] = c
			options[c] = map[int]bool{7: true}
		case 4:
			answers[4] = c
			options[c] = map[int]bool{4: true}
		case 5:
			options[c] = map[int]bool{
				2: true,
				3: true,
				5: true,
			}
		case 6:
			options[c] = map[int]bool{
				0: true,
				6: true,
				9: true,
			}
		case 7:
			answers[8] = c
			options[c] = map[int]bool{8: true}
		default:
			panic("bad input")
		}
	}

	for _, c := range inputs {
		if len(options[c]) == 1 {
			continue
		}
		if contains(c, answers[4]) && contains(c, answers[7]) {
			answers[9] = c
			options[c] = map[int]bool{9: true}
		} else {
			delete(options[c], 9)
		}
	}

	// narrow down based on knowledge of 1
	for _, c := range inputs {
		if len(options[c]) == 1 {
			continue
		}
		if !contains(c, answers[1]) {
			delete(options[c], 3)
			delete(options[c], 4)
			delete(options[c], 9)
		}
		if !contains(c, answers[4]) {
			delete(options[c], 9)
		}
		if !contains(c, answers[7]) {
			delete(options[c], 9)
			delete(options[c], 0)
			delete(options[c], 3)
		}
		if contains(answers[9], c) {
			delete(options[c], 6)
			delete(options[c], 2)
			delete(options[c], 0)
		}
		if len(options[c]) == 1 && options[c][5] {
			answers[5] = c
		}
	}

	if _, ok := answers[5]; !ok {
		panic("flawed")
	}
	for c, o := range options {
		if c != answers[5] {
			delete(o, 5)
		}
		if len(options[c]) == 1 && options[c][6] {
			answers[6] = c
		}
	}

	if _, ok := answers[6]; !ok {
		panic("flawed")
	}

	for c, o := range options {
		if c != answers[6] {
			delete(o, 6)
		}
		if len(options[c]) == 1 && options[c][2] {
			answers[2] = c
		}
		if len(options[c]) == 1 && options[c][3] {
			answers[3] = c
		}
		if len(options[c]) == 1 && options[c][0] {
			answers[0] = c
		}
	}

	out := 0
	for _, o := range outputs {
		out *= 10
		out += only(options[o])
	}
	return out
}

func contains(s, c string) bool {
	for i := range c {
		if !strings.Contains(s, c[i:i+1]) {
			return false
		}
	}
	return true
}

func sortItems(in []string) (out []string) {
	for _, s := range in {
		out = append(out, sortString(s))
	}
	return out
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

func only(m map[int]bool) int {
	for x := range m {
		return x
	}
	panic("bad")
}
