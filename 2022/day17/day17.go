package day17

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
Part1 Prompt

--- Day 17: Pyroclastic Flow ---
Your handheld device has located an alternative exit from the cave for you and
the elephants. The ground is rumbling almost continuously now, but the strange
valves bought you some time. It's definitely getting warmer in here, though.

The tunnels eventually open into a very tall, narrow chamber. Large,
oddly-shaped rocks are falling into the chamber from above, presumably due to
all the rumbling. If you can't work out where the rocks will fall next, you
might be crushed!

The five types of rocks have the following peculiar shapes, where # is rock and
. is empty space:

	####

	.#.
	###
	.#.

	..#
	..#
	###

	#
	#
	#
	#

	##
	##

The rocks fall in the order shown above: first the - shape, then the + shape,
and so on. Once the end of the list is reached, the same order repeats: the -
shape falls first, sixth, 11th, 16th, etc.

The rocks don't spin, but they do get pushed around by jets of hot gas coming
out of the walls themselves. A quick scan reveals the effect the jets of hot
gas will have on the rocks as they fall (your puzzle input).

For example, suppose this was the jet pattern in your cave:

	>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>

In jet patterns, < means a push to the left, while > means a push to the right.
The pattern above means that the jets will push a falling rock right, then
right, then right, then left, then left, then right, and so on. If the end of
the list is reached, it repeats.

The tall, vertical chamber is exactly seven units wide. Each rock appears so
that its left edge is two units away from the left wall and its bottom edge is
three units above the highest rock in the room (or the floor, if there isn't
one).

After a rock appears, it alternates between being pushed by a jet of hot gas
one unit (in the direction indicated by the next symbol in the jet pattern) and
then falling one unit down. If any movement would cause any part of the rock to
move into the walls, floor, or a stopped rock, the movement instead does not
occur. If a downward movement would have caused a falling rock to move into the
floor or an already-fallen rock, the falling rock stops where it is (having
landed on something) and a new rock immediately begins falling.

Drawing falling rocks with @ and stopped rocks with #, the jet pattern in the
example above manifests as follows:

	The first rock begins falling:
	|..@@@@.|
	|.......|
	|.......|
	|.......|
	+-------+

	Jet of gas pushes rock right:
	|...@@@@|
	|.......|
	|.......|
	|.......|
	+-------+

	Rock falls 1 unit:
	|...@@@@|
	|.......|
	|.......|
	+-------+

	Jet of gas pushes rock right, but nothing happens:
	|...@@@@|
	|.......|
	|.......|
	+-------+

	Rock falls 1 unit:
	|...@@@@|
	|.......|
	+-------+

	Jet of gas pushes rock right, but nothing happens:
	|...@@@@|
	|.......|
	+-------+

	Rock falls 1 unit:
	|...@@@@|
	+-------+

	Jet of gas pushes rock left:
	|..@@@@.|
	+-------+

	Rock falls 1 unit, causing it to come to rest:
	|..####.|
	+-------+

	A new rock begins falling:
	|...@...|
	|..@@@..|
	|...@...|
	|.......|
	|.......|
	|.......|
	|..####.|
	+-------+

	Jet of gas pushes rock left:
	|..@....|
	|.@@@...|
	|..@....|
	|.......|
	|.......|
	|.......|
	|..####.|
	+-------+

	Rock falls 1 unit:
	|..@....|
	|.@@@...|
	|..@....|
	|.......|
	|.......|
	|..####.|
	+-------+

	Jet of gas pushes rock right:
	|...@...|
	|..@@@..|
	|...@...|
	|.......|
	|.......|
	|..####.|
	+-------+

	Rock falls 1 unit:
	|...@...|
	|..@@@..|
	|...@...|
	|.......|
	|..####.|
	+-------+

	Jet of gas pushes rock left:
	|..@....|
	|.@@@...|
	|..@....|
	|.......|
	|..####.|
	+-------+

	Rock falls 1 unit:
	|..@....|
	|.@@@...|
	|..@....|
	|..####.|
	+-------+

	Jet of gas pushes rock right:
	|...@...|
	|..@@@..|
	|...@...|
	|..####.|
	+-------+

	Rock falls 1 unit, causing it to come to rest:
	|...#...|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	A new rock begins falling:
	|....@..|
	|....@..|
	|..@@@..|
	|.......|
	|.......|
	|.......|
	|...#...|
	|..###..|
	|...#...|
	|..####.|
	+-------+

The moment each of the next few rocks begins falling, you would see this:

	|..@....|
	|..@....|
	|..@....|
	|..@....|
	|.......|
	|.......|
	|.......|
	|..#....|
	|..#....|
	|####...|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|..@@...|
	|..@@...|
	|.......|
	|.......|
	|.......|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|..@@@@.|
	|.......|
	|.......|
	|.......|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|...@...|
	|..@@@..|
	|...@...|
	|.......|
	|.......|
	|.......|
	|.####..|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|....@..|
	|....@..|
	|..@@@..|
	|.......|
	|.......|
	|.......|
	|..#....|
	|.###...|
	|..#....|
	|.####..|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|..@....|
	|..@....|
	|..@....|
	|..@....|
	|.......|
	|.......|
	|.......|
	|.....#.|
	|.....#.|
	|..####.|
	|.###...|
	|..#....|
	|.####..|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|..@@...|
	|..@@...|
	|.......|
	|.......|
	|.......|
	|....#..|
	|....#..|
	|....##.|
	|....##.|
	|..####.|
	|.###...|
	|..#....|
	|.####..|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

	|..@@@@.|
	|.......|
	|.......|
	|.......|
	|....#..|
	|....#..|
	|....##.|
	|##..##.|
	|######.|
	|.###...|
	|..#....|
	|.####..|
	|....##.|
	|....##.|
	|....#..|
	|..#.#..|
	|..#.#..|
	|#####..|
	|..###..|
	|...#...|
	|..####.|
	+-------+

To prove to the elephants your simulation is accurate, they want to know how
tall the tower will get after 2022 rocks have stopped (but before the 2023rd
rock begins falling). In this example, the tower of rocks will be 3068 units
tall.

How many units tall will the tower of rocks be after 2022 rocks have stopped
falling?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day17(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day17(r)
}

const numRocks = 2022

func day17(r io.Reader) (answer int, err error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	moves := string(bytes.TrimSpace(b))

	g := newGrid()
	maxHeight := 0
	j := 0

	for i := 0; i < numRocks; i++ {
		b := blocks[i%len(blocks)]
		b.SetLoc(2, maxHeight+3+1) // want a gap of 3
		for {
			debug(g, b)
			var dx int
			if moves[j%len(moves)] == '>' {
				dx = 1
			} else {
				dx = -1
			}
			if b.CanMove(g, dx, 0) {
				b.Move(dx, 0)
			}
			j++

			debug(g, b)
			if !b.CanMove(g, 0, -1) {
				g.Place(b)
				maxHeight = max(maxHeight, b.MaxY())
				break
			}
			b.Move(0, -1)
		}
	}
	return maxHeight, nil
}

func debug(g *grid, b *block) {
	if os.Getenv("DEBUG") == "" {
		return
	}
	g.Print(b)
	_, _ = fmt.Scanln()
}

type block struct {
	x, y          int
	cells         [][2]int
	width, height int
}

func (b *block) SetLoc(x int, y int) {
	b.x, b.y = x, y
}

func (b *block) TryMoveX(g *grid, dx int) {
	if newX := b.x + dx; newX > 0 && newX+b.width <= 7 {
		b.x = newX
	}
}

func (b *block) CanMove(g *grid, dx, dy int) bool {
	newX, newY := b.x+dx, b.y+dy
	if newX < 0 || newX+b.width > 7 {
		return false
	}
	if newY <= 0 {
		return false
	}
	for _, c := range b.cells {
		x, y := newX+c[0], newY+c[1]
		if g.Has(x, y) {
			return false
		}
	}
	return true
}

func (b *block) Move(dx, dy int) {
	b.x, b.y = b.x+dx, b.y+dy
}

func (b *block) MaxY() int {
	return b.y + b.height - 1
}

var blocks = []*block{
	{
		cells: [][2]int{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		width:  4,
		height: 1,
	},
	{
		cells: [][2]int{
			{1, 0},
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 2},
		},
		width:  3,
		height: 3,
	},
	{
		cells: [][2]int{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
		},
		width:  3,
		height: 3,
	},
	{
		cells: [][2]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
		width:  1,
		height: 4,
	},
	{
		cells: [][2]int{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		},
		width:  2,
		height: 2,
	},
}

type grid struct {
	filled map[[2]int]bool
}

func newGrid() *grid {
	return &grid{
		filled: map[[2]int]bool{},
	}
}

func (g *grid) Place(b *block) {
	for _, c := range b.cells {
		g.filled[[2]int{b.x + c[0], b.y + c[1]}] = true
	}
}

func (g *grid) Has(x int, y int) bool {
	return g.filled[[2]int{x, y}]
}

func (g *grid) Print(b *block) {
	fmt.Println()
	bHas := map[[2]int]bool{}
	for _, c := range b.cells {
		x, y := b.x+c[0], b.y+c[1]
		bHas[[2]int{x, y}] = true
	}
	for y := b.MaxY(); y > 0; y-- {
		for x := 0; x < 7; x++ {
			l := [2]int{x, y}
			if bHas[l] {
				fmt.Print("@")
			} else if g.Has(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("-------")
}
