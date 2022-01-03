package day24

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (answer int, err error) {
	return day24(r, true)
}

func day24(r io.Reader, max bool) (answer int, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	const numBlocks = 14
	linesPerBlock := len(lines) / numBlocks
	var stack []int
	var digits [numBlocks]int
	digits[0] = 1
	for i := 0; i < numBlocks; i++ {
		if lines[(i*linesPerBlock)+4] == "div z 1" {
			stack = append(stack, i)
		} else {
			leftI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			leftIncr := num(strings.Split(lines[leftI*linesPerBlock+15], " ")[2])
			rightIncr := num(strings.Split(lines[i*linesPerBlock+5], " ")[2])
			digits[leftI], digits[i] = match(leftIncr, rightIncr, max)
		}
	}
	for i := 0; i < numBlocks; i++ {
		answer *= 10
		answer += digits[i]
	}
	return answer, err
}

func match(left, right int, max bool) (int, int) {
	if max {
		if left+right <= 0 {
			return 9, 9 + left + right
		}
		return 9 - left - right, 9
	}
	if
}

func num(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}
