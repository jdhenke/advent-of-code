package day14

import (
	"bufio"
	"io"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
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

	for i := 0; i < 10; i++ {
		type replacement struct {
			i int
			c string
		}
		var replacements []replacement
		for i := 0; i+1 < len(text); i++ {
			if c, ok := rules[text[i:i+2]]; ok {
				replacements = append(replacements, replacement{
					i: i,
					c: c,
				})
			}
		}
		for i := 0; i < len(replacements); i++ {
			r := replacements[len(replacements)-i-1]
			text = text[:r.i+1] + r.c + text[r.i+1:]
		}
	}

	var min, max string
	freqs := make(map[string]int)
	for i := 0; i < len(text); i++ {
		freqs[text[i:i+1]]++
	}
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
