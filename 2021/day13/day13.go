package day13

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type entry struct {
	i, j int
}

var re = regexp.MustCompile(`fold along ([xy])=(\d+)`)

func Part1(r io.Reader) (ans int, err error) {
	return day11(r, 1)
}

func Part2(r io.Reader) (ans int, err error) {
	return day11(r, 0)
}

func day11(r io.Reader, numFolds int) (ans int, err error) {
	dots := make(map[entry]bool)
	scan := bufio.NewScanner(r)
	scan.Scan()
	maxX, maxY := 0, 0
	for scan.Text() != "" {
		parts := strings.Split(scan.Text(), ",")
		xs, ys := parts[0], parts[1]
		x, err := strconv.Atoi(xs)
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(ys)
		if err != nil {
			return 0, err
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		dots[entry{x, y}] = true
		scan.Scan()
	}
	folds := 0
	for scan.Scan() && (numFolds == 0 || folds < numFolds) {
		folds++
		// fold along y=7
		m := re.FindStringSubmatch(scan.Text())
		direction := m[1]
		num, err := strconv.Atoi(m[2])
		if err != nil {
			return 0, err
		}
		if m == nil {
			return 0, fmt.Errorf("bad line: %v", scan.Text())
		}
		switch direction {
		case "x":
			for x := num + 1; x <= maxX; x++ {
				for y := 0; y <= maxY; y++ {
					if dots[entry{x, y}] {
						dots[entry{num - (x - num), y}] = true
						delete(dots, entry{x, y})
					}
				}
			}
			maxX = num
		case "y":
			for y := num + 1; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if dots[entry{x, y}] {
						dots[entry{x, num - (y - num)}] = true
						delete(dots, entry{x, y})
					}
				}
			}
			maxY = num
		}
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}
	if numFolds == 0 {
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				c := " "
				if dots[entry{x, y}] {
					c = "#"
				}
				fmt.Printf(c)
			}
			fmt.Printf("\n")
		}
	}
	return len(dots), nil
}
