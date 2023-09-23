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

/*
Part1 Prompt

--- Day 13: Transparent Origami ---
You reach another volcanically active part of the cave. It would be nice if you
could do some kind of thermal imaging so you could tell ahead of time which
caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you
activate it, you are greeted with:

	Congratulations on your purchase! To activate this infrared thermal imaging
	camera system, please enter the code found on page 1 of the manual.

Apparently, the Elves have never used this feature. To your surprise, you
manage to find the manual; as you go to open it, page 1 falls out. It's a large
sheet of transparent paper! The transparent paper is marked with random dots
and includes instructions on how to fold it up (your puzzle input). For
example:

	6,10
	0,14
	9,10
	0,3
	10,4
	4,11
	6,0
	6,12
	4,1
	0,13
	10,12
	3,4
	3,0
	8,4
	1,10
	2,14
	8,10
	9,0

	fold along y=7
	fold along x=5

The first section is a list of dots on the transparent paper. 0,0 represents
the top-left coordinate. The first value, x, increases to the right. The second
value, y, increases downward. So, the coordinate 3,0 is to the right of 0,0,
and the coordinate 0,7 is below 0,0. The coordinates in this example form the
following pattern, where # is a dot on the paper and . is an empty, unmarked
position:

	...#..#..#.
	....#......
	...........
	#..........
	...#....#.#
	...........
	...........
	...........
	...........
	...........
	.#....#.##.
	....#......
	......#...#
	#..........
	#.#........

Then, there is a list of fold instructions. Each instruction indicates a line
on the transparent paper and wants you to fold the paper up (for horizontal
y=... lines) or left (for vertical x=... lines). In this example, the first
fold instruction is fold along y=7, which designates the line formed by all of
the positions where y is 7 (marked here with -):

	...#..#..#.
	....#......
	...........
	#..........
	...#....#.#
	...........
	...........
	-----------
	...........
	...........
	.#....#.##.
	....#......
	......#...#
	#..........
	#.#........

Because this is a horizontal line, fold the bottom half up. Some of the dots
might end up overlapping after the fold is complete, but dots will never appear
exactly on a fold line. The result of doing this fold looks like this:

	#.##..#..#.
	#...#......
	......#...#
	#...#......
	.#.#..#.###
	...........
	...........

Now, only 17 dots are visible.

Notice, for example, the two dots in the bottom left corner before the
transparent paper is folded; after the fold is complete, those dots appear in
the top left corner (at 0,0 and 0,1). Because the paper is transparent, the dot
just below them in the result (at 0,3) remains visible, as it can be seen
through the transparent paper.

Also notice that some dots can end up overlapping; in this case, the dots merge
together and become a single dot.

The second fold instruction is fold along x=5, which indicates this line:

	#.##.|#..#.
	#...#|.....
	.....|#...#
	#...#|.....
	.#.#.|#.###
	.....|.....
	.....|.....

Because this is a vertical line, fold left:

	#####
	#...#
	#...#
	#...#
	#####
	.....
	.....

The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the
first fold. After the first fold in the example above, 17 dots are visible -
dots that end up overlapping after the fold is completed count as a single dot.

How many dots are visible after completing just the first fold instruction on
your transparent paper?
*/
func Part1(r io.Reader) (ans int, err error) {
	return day11(r, 1)
}

/*
Part2 Prompt

--- Part Two ---
Finish folding the transparent paper according to the instructions. The manual
says the code is always eight capital letters.

What code do you use to activate the infrared thermal imaging camera system?
*/
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
				fmt.Print(c)
			}
			fmt.Printf("\n")
		}
	}
	return len(dots), nil
}
