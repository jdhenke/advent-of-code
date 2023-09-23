package day12

import (
	"io"
	"strconv"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 12: Rain Risk ---
Your ferry made decent progress toward the island, but the storm came in faster
than anyone expected. The ferry needs to take evasive actions!

Unfortunately, the ship's navigation computer seems to be malfunctioning;
rather than giving a route directly to safety, it produced extremely circuitous
instructions. When the captain uses the PA system to ask if anyone can help,
you quickly volunteer.

The navigation instructions (your puzzle input) consists of a sequence of
single-character actions paired with integer input values. After staring at
them for a few minutes, you work out what they probably mean:

- Action N means to move north by the given value.
- Action S means to move south by the given value.
- Action E means to move east by the given value.
- Action W means to move west by the given value.
- Action L means to turn left the given number of degrees.
- Action R means to turn right the given number of degrees.
- Action F means to move forward by the given value in the direction the ship
is currently facing.

The ship starts by facing east. Only the L and R actions change the direction
the ship is facing. (That is, if the ship is facing east and the next
instruction is N10, the ship would move north 10 units, but would still move
east if the following action were F.)

For example:

	F10
	N3
	F7
	R90
	F11

These instructions would be handled as follows:

- F10 would move the ship 10 units east (because the ship starts by facing
east) to east 10, north 0.
- N3 would move the ship 3 units north to east 10, north 3.
- F7 would move the ship another 7 units east (because the ship is still facing
east) to east 17, north 3.
- R90 would cause the ship to turn right by 90 degrees and face south; it
remains at east 17, north 3.
- F11 would move the ship 11 units south to east 17, south 8.

At the end of these instructions, the ship's Manhattan distance (sum of the
absolute values of its east/west position and its north/south position) from
its starting position is 17 + 8 = 25.

Figure out where the navigation instructions lead. What is the Manhattan
distance between that location and the ship's starting position?
*/
func Part1(r io.Reader) (answer int, err error) {
	dir := east
	return day12(r, func(c string, n int) vector {
		switch c {
		case "N", "S", "E", "W", "F", "B":
			d := map[string]vector{
				"N": north,
				"S": south,
				"E": east,
				"W": west,
				"F": dir,
				"B": dir.backwards(),
			}[c]
			return d.scale(n)
		case "L", "R":
			dir = dir.rotate(c, n)
			return vector{}
		default:
			panic("bad command: " + c)
		}
	})
}

/*
Part2 Prompt

--- Part Two ---
Before you can give the destination to the captain, you realize that the actual
action meanings were printed on the back of the instructions the whole time.

Almost all of the actions indicate how to move a waypoint which is relative to
the ship's position:

- Action N means to move the waypoint north by the given value.
- Action S means to move the waypoint south by the given value.
- Action E means to move the waypoint east by the given value.
- Action W means to move the waypoint west by the given value.
- Action L means to rotate the waypoint around the ship left
(counter-clockwise) the given number of degrees.
- Action R means to rotate the waypoint around the ship right (clockwise) the
given number of degrees.
- Action F means to move forward to the waypoint a number of times equal to the
given value.

The waypoint starts 10 units east and 1 unit north relative to the ship. The
waypoint is relative to the ship; that is, if the ship moves, the waypoint
moves with it.

For example, using the same instructions as above:

- F10 moves the ship to the waypoint 10 times (a total of 100 units east and 10
units north), leaving the ship at east 100, north 10. The waypoint stays 10
units east and 1 unit north of the ship.
- N3 moves the waypoint 3 units north to 10 units east and 4 units north of the
ship. The ship remains at east 100, north 10.
- F7 moves the ship to the waypoint 7 times (a total of 70 units east and 28
units north), leaving the ship at east 170, north 38. The waypoint stays 10
units east and 4 units north of the ship.
- R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4
units east and 10 units south of the ship. The ship remains at east 170, north
38.
- F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110
units south), leaving the ship at east 214, south 72. The waypoint stays 4
units east and 10 units south of the ship.

After these operations, the ship's Manhattan distance from its starting
position is 214 + 72 = 286.

Figure out where the navigation instructions actually lead. What is the
Manhattan distance between that location and the ship's starting position?
*/
func Part2(r io.Reader) (answer int, err error) {
	wp := vector{10, 1}
	return day12(r, func(c string, n int) vector {
		switch c {
		case "F", "B":
			d := wp
			if c == "B" {
				d = d.backwards()
			}
			return d.scale(n)
		case "L", "R":
			wp = wp.rotate(c, n)
			return vector{}
		case "N", "S", "E", "W":
			d := map[string]vector{
				"N": north,
				"S": south,
				"E": east,
				"W": west,
			}[c]
			wp = wp.add(d.scale(n))
			return vector{}
		default:
			panic("bad command:" + c)
		}
	})
}

func day12(r io.Reader, move func(c string, n int) vector) (answer int, err error) {
	loc := vector{}
	if err := input.ForEachLine(r, func(line string) error {
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		c := line[:1]
		loc = loc.add(move(c, n))
		return nil
	}); err != nil {
		return 0, err
	}
	return loc.manhattan(), nil
}

var (
	north = vector{0, 1}
	east  = vector{1, 0}
	south = vector{0, -1}
	west  = vector{-1, 0}
)

type vector struct {
	dx, dy int
}

func (d vector) manhattan() int {
	return abs(d.dx) + abs(d.dy)
}

func (d vector) scale(n int) vector {
	return vector{n * d.dx, n * d.dy}
}

func (d vector) add(other vector) vector {
	return vector{d.dx + other.dx, d.dy + other.dy}
}

func (d vector) backwards() vector {
	return vector{-d.dx, -d.dy}
}

func (d vector) rotate(c string, n int) vector {
	n = n % 360
	if c == "L" {
		n = 360 - n
	}
	for i := 0; i < n/90; i++ {
		d.dx, d.dy = d.dy, -d.dx
	}
	return d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
