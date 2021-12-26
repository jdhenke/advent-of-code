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
	covered := map[string]bool{"start": true}
	current := "start"
	return numPaths(g, covered, current), nil
}

func numPaths(g map[string][]string, covered map[string]bool, current string) int {
	total := 0
	for _, n := range g[current] {
		if covered[n] {
			continue
		}
		if n == "end" {
			total += 1
		} else {
			if strings.ToLower(n) == n {
				covered[n] = true
			}
			total += numPaths(g, covered, n)
			delete(covered, n)
		}
	}
	return total
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
