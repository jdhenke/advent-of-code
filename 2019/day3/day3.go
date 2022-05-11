package day3

import (
	"io"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 3: Crossed Wires ---
The gravity assist was successful, and you're well on your way to the Venus
refuelling station. During the rush back on Earth, the fuel management system
wasn't completely installed, so that's next on the priority list.

Opening the front panel reveals a jumble of wires. Specifically, two wires are
connected to a central port and extend outward on a grid. You trace the path
each wire takes as it leaves the central port, one wire per line of text (your
puzzle input).

The wires twist and turn, but the two wires occasionally cross paths. To fix
the circuit, you need to find the intersection point closest to the central
port. Because the wires are on a grid, use the Manhattan distance for this
measurement. While the wires do technically cross right at the central port
where they both start, this point does not count, nor does a wire count as
crossing with itself.

For example, if the first wire's path is R8,U5,L5,D3, then starting from the
central port (o), it goes right 8, up 5, left 5, and finally down 3:

    ...........
    ...........
    ...........
    ....+----+.
    ....|....|.
    ....|....|.
    ....|....|.
    .........|.
    .o-------+.
    ...........

Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6, down 4,
and left 4:

    ...........
    .+-----+...
    .|.....|...
    .|..+--X-+.
    .|..|..|.|.
    .|.-X--+.|.
    .|..|....|.
    .|.......|.
    .o-------+.
    ...........

These wires cross at two locations (marked X), but the lower-left one is closer
to the central port: its distance is 3 + 3 = 6.

Here are a few more examples:

- R75,D30,R83,U83,L12,D49,R71,U7,L72U62,R66,U55,R34,D71,R55,D58,R83 = distance
159
-
R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51U98,R91,D20,R16,D67,R40,U7,R15,U6,R7
= distance 135

What is the Manhattan distance from the central port to the closest
intersection?
*/
func Part1(r io.Reader) (answer int, err error) {
	if err := day3(r, func(p point) {
		manhattan := abs(p.x) + abs(p.y)
		if answer == 0 || manhattan < answer {
			answer = manhattan
		}
	}); err != nil {
		return 0, err
	}
	return answer, nil
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}

type point struct {
	x, y int
}

// call fn for each intersection
func day3(r io.Reader, fn func(p point)) error {
	var pointsPerLine []map[point]struct{}
	if err := input.ForEachLine(r, func(line string) error {
		points := make(map[point]struct{})
		x, y := 0, 0
		steps := strings.Split(line, ",")
		for _, step := range steps {
			d := step[:1]
			n, err := strconv.Atoi(step[1:])
			if err != nil {
				return err
			}
			var move func()
			switch d {
			case "U":
				move = func() { y++ }
			case "D":
				move = func() { y-- }
			case "R":
				move = func() { x++ }
			case "L":
				move = func() { x-- }
			default:
				panic("unknown direction: " + d)
			}
			for i := 0; i < n; i++ {
				move()
				points[point{x, y}] = struct{}{}
			}
		}
		pointsPerLine = append(pointsPerLine, points)
		return nil
	}); err != nil {
		return err
	}
	for p := range pointsPerLine[0] {
		if _, ok := pointsPerLine[1][p]; ok {
			fn(p)
		}
	}
	return nil
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
