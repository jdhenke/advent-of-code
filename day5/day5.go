package day5

import (
	"advent-of-code/input"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func Part1(r io.Reader) (ans int, err error) {
	return day5(r, false)
}

func Part2(r io.Reader) (ans int, err error) {
	return day5(r, true)
}

func day5(r io.Reader, diagonal bool) (int, error) {
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	type entry struct {
		x, j int
	}
	spots := make(map[entry]int)
	if err := input.ForEachLine(r, func(line string) error {
		match := re.FindStringSubmatch(line)
		if match == nil {
			return fmt.Errorf("bad line: %v", line)
		}
		x1, y1, x2, y2 := mustInt(match[1]), mustInt(match[2]), mustInt(match[3]), mustInt(match[4])
		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				spots[entry{x1, y}]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				spots[entry{x, y1}]++
			}
		} else if diagonal {
			dx, dy := 1, 1
			if x2 < x1 {
				dx = -1
			}
			if y2 < y1 {
				dy = -1
			}
			for x, y := x1, y1; x != x2 && y != y2; x, y = x+dx, y+dy {
				spots[entry{x, y}]++
			}
			spots[entry{x2, y2}]++
		}
		return nil
	}); err != nil {
		return 0, err
	}
	count := 0
	for _, val := range spots {
		if val >= 2 {
			count++
		}
	}
	return count, nil
}

func mustInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
