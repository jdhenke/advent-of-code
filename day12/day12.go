package day12

import (
	"advent-of-code/input"
	"io"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	g, err := getGraph(r)
	if err != nil {
		return 0, err
	}
	current := "start"
	return numPaths(g, nil, current, func(stack []string, next string) bool {
		return strings.ToLower(next) != next || !in(stack, next)
	}), nil
}

func Part2(r io.Reader) (ans int, err error) {
	g, err := getGraph(r)
	if err != nil {
		return 0, err
	}
	current := "start"
	return numPaths(g, nil, current, func(stack []string, next string) bool {
		if strings.ToLower(next) != next {
			return true
		}
		freqs := make(map[string]int)
		hasDouble := false
		for _, s := range stack {
			if strings.ToLower(s) != s {
				continue
			}
			freqs[s]++
			if freqs[s] > 1 {
				hasDouble = true
			}
		}
		if freqs[next] == 0 {
			return true
		}
		if freqs[next] == 1 && !hasDouble {
			return true
		}
		return false
	}), nil
}

func numPaths(g map[string][]string, stack []string, current string, allow func(stack []string, next string) bool) int {
	total := 0
	for _, n := range g[current] {
		if n == "start" {
			continue
		}
		if n == "end" {
			total ++
			//fmt.Println(stack)
			continue
		}
		if !allow(stack, n) {
			continue
		}
		stack = append(stack, n)
		total += numPaths(g, stack, n, allow)
		stack = stack[:len(stack)-1]
	}
	return total
}

func in(stack []string, s string) bool {
	for _, x := range stack {
		if x == s {
			return true
		}
	}
	return false
}

func getGraph(r io.Reader) (map[string][]string, error) {
	m := make(map[string][]string)
	if err := input.ForEachLine(r, func(line string) error {
		nodes := strings.Split(line, "-")
		n1, n2 := nodes[0], nodes[1]
		m[n1] = append(m[n1], n2)
		m[n2] = append(m[n2], n1)
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}
