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
	numFolds := 1
	folds := 0
	for scan.Scan() && folds < numFolds {
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
			maxX = num - 1
		case "y":
			for y := num + 1; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if dots[entry{x, y}] {
						dots[entry{x, num - (y - num)}] = true
						delete(dots, entry{x, y})
					}
				}
			}
			maxY = num - 1
		}
	}
	if err := scan.Err(); err != nil {
		return 0, err
	}
	return len(dots), nil
}
