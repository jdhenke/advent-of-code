package day1

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

const size = 3

func Part1(r io.Reader) (answer string, err error) {
	s := bufio.NewScanner(r)
	var last = math.MaxInt
	var increases int
	for s.Scan() {
		text := s.Text()
		n, err := strconv.Atoi(text)
		if err != nil {
			return "", fmt.Errorf("error converting '%s' to int: %v", text, err)
		}
		if n > last {
			increases++
		}
		last = n
	}
	if err := s.Err(); err != nil {
		return "", fmt.Errorf("error scanning text: %v", err)
	}
	return fmt.Sprint(increases), nil
}

func Part2(r io.Reader) (answer string, err error) {
	s := bufio.NewScanner(r)
	var recent [size]int
	var increases int
	i := 0
	sum := 0
	var last int
	for s.Scan() {
		text := s.Text()
		n, err := strconv.Atoi(text)
		if err != nil {
			return "", fmt.Errorf("error converting '%s' to int: %v", text, err)
		}
		i++
		if i > 3 {
			sum -= recent[i%3]
		}
		sum += n
		recent[i%3] = n
		if i > 3 {
			if sum > last {
				increases++
			}
		}
		last = sum
	}
	if err := s.Err(); err != nil {
		return "", fmt.Errorf("error scanning text: %v", err)
	}
	return fmt.Sprint(increases), nil
}
