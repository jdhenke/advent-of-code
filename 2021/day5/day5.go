package day5

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents
constantly produce large, opaque clouds, so it would be best to avoid them if
possible.

They tend to form in lines; the submarine helpfully produces a list of nearby
lines of vents (your puzzle input) for you to review. For example:

	0,9 -> 5,9
	8,0 -> 0,8
	9,4 -> 3,4
	2,2 -> 2,1
	7,0 -> 7,4
	6,4 -> 2,0
	0,9 -> 2,9
	3,4 -> 1,4
	0,0 -> 8,8
	5,5 -> 8,2

Each line of vents is given as a line segment in the format x1,y1 -> x2,y2
where x1,y1 are the coordinates of one end the line segment and x2,y2 are the
coordinates of the other end. These line segments include the points at both
ends. In other words:

- An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
- An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.

For now, only consider horizontal and vertical lines: lines where either x1 =
x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the
following diagram:

	.......1..
	..1....1..
	..1....1..
	.......1..
	.112111211
	..........
	..........
	..........
	..........
	222111....

In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9.
Each position is shown as the number of lines which cover that point or . if no
line covers that point. The top-left pair of 1s, for example, comes from 2,2 ->
2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9
-> 2,9.

To avoid the most dangerous areas, you need to determine the number of points
where at least two lines overlap. In the above example, this is anywhere in the
diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two
lines overlap?
*/
func Part1(r io.Reader) (ans int, err error) {
	return day5(r, false)
}

/*
Part2 Prompt

--- Part Two ---
Unfortunately, considering only horizontal and vertical lines doesn't give you
the full picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in
your list will only ever be horizontal, vertical, or a diagonal line at exactly
45 degrees. In other words:

- An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
- An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.

Considering all lines from the above example would now produce the following
diagram:

	1.1....11.
	.111...2..
	..2.1.111.
	...1.2.2..
	.112313211
	...1.2....
	..1...1...
	.1.....1..
	1.......1.
	222111....

You still need to determine the number of points where at least two lines
overlap. In the above example, this is still anywhere in the diagram with a 2
or larger - now a total of 12 points.

Consider all of the lines. At how many points do at least two lines overlap?
*/
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
		if !diagonal && !((x1 == x2) || (y1 == y2)) {
			return nil
		}
		dx, dy := diff(x1, x2), diff(y1, y2)
		for x, y := x1, y1; !(x == x2 && y == y2); x, y = x+dx, y+dy {
			spots[entry{x, y}]++
		}
		spots[entry{x2, y2}]++
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

func diff(x1, x2 int) int {
	if x1 == x2 {
		return 0
	}
	if x1 < x2 {
		return 1
	}
	return -1
}
