package day14

import (
	"bufio"
	"io"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	return day14(r, 10)
}

func Part2(r io.Reader) (ans int, err error) {
	return day14(r, 40)
}

func day14(r io.Reader, steps int) (ans int, err error) {
	scan := bufio.NewScanner(r)
	scan.Scan()
	text := scan.Text()
	scan.Scan()
	rules := make(map[string]string)
	for scan.Scan() {
		parts := strings.Split(scan.Text(), " -> ")
		rules[parts[0]] = parts[1]
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}

	type memoEntry struct {
		text   string
		rounds int
	}
	memo := make(map[memoEntry]map[string]int)
	var getCounts func(text string, rounds int) (counts map[string]int)
	getCounts = func(text string, rounds int) (counts map[string]int) {
		e := memoEntry{
			text:   text,
			rounds: rounds,
		}
		if counts, ok := memo[e]; ok {
			return counts
		}
		defer func() {
			memo[e] = counts
		}()
		if rounds == 0 {
			return nil
		}
		counts = map[string]int{}
		for i := 0; i+1 < len(text); i++ {
			if m, ok := rules[text[i:i+2]]; ok {
				counts[m]++
				for c, v := range getCounts(text[i:i+1]+m, rounds-1) {
					counts[c] += v
				}
				for c, v := range getCounts(m+text[i+1:i+2], rounds-1) {
					counts[c] += v
				}
			}
		}
		return counts
	}
	freqs := getCounts(text, steps)
	for i := 0; i < len(text); i++ {
		freqs[text[i:i+1]]++
	}
	var min, max string
	for c := range freqs {
		if min == "" || freqs[c] < freqs[min] {
			min = c
		}
		if max == "" || freqs[c] > freqs[max] {
			max = c
		}
	}
	return freqs[max] - freqs[min], nil
}
