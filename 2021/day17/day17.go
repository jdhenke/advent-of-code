package day17

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

/*
Part1 Prompt

--- Day 17: Trick Shot ---
You finally decode the Elves' message. HI, the message says. You continue
searching for the sleigh keys.

Ahead of you is what appears to be a large ocean trench. Could the keys have
fallen into it? You'd better send a probe to investigate.

The probe launcher on your submarine can fire the probe with any integer
velocity in the x (forward) and y (upward, or downward if negative) directions.
For example, an initial x,y velocity like 0,10 would fire the probe straight
up, while an initial velocity like 10,-1 would fire the probe forward at a
slight downward angle.

The probe's x,y position starts at 0,0. Then, it will follow some trajectory by
moving in steps. On each step, these changes occur in the following order:

- The probe's x position increases by its x velocity.
- The probe's y position increases by its y velocity.
- Due to drag, the probe's x velocity changes by 1 toward the value 0; that is,
it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0,
or does not change if it is already 0.
- Due to gravity, the probe's y velocity decreases by 1.

For the probe to successfully make it into the trench, the probe must be on
some trajectory that causes it to be within a target area after any step. The
submarine computer has already calculated this target area (your puzzle input).
For example:

	target area: x=20..30, y=-10..-5

This target area means that you need to find initial x,y velocity values such
that after any step, the probe's x position is at least 20 and at most 30, and
the probe's y position is at least -10 and at most -5.

Given this target area, one initial velocity that causes the probe to be within
the target area after any step is 7,2:

	.............#....#............
	.......#..............#........
	...............................
	S........................#.....
	...............................
	...............................
	...........................#...
	...............................
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTT#TT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT

In this diagram, S is the probe's initial position, 0,0. The x coordinate
increases to the right, and the y coordinate increases upward. In the bottom
right, positions that are within the target area are shown as T. After each
step (until the target area is reached), the position of the probe is marked
with #. (The bottom-right # is both a position the probe reaches and a position
in the target area.)

Another initial velocity that causes the probe to be within the target area
after any step is 6,3:

	...............#..#............
	...........#........#..........
	...............................
	......#..............#.........
	...............................
	...............................
	S....................#.........
	...............................
	...............................
	...............................
	.....................#.........
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................T#TTTTTTTTT
	....................TTTTTTTTTTT

Another one is 9,0:

	S........#.....................
	.................#.............
	...............................
	........................#......
	...............................
	....................TTTTTTTTTTT
	....................TTTTTTTTTT#
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT
	....................TTTTTTTTTTT

One initial velocity that doesn't cause the probe to be within the target area
after any step is 17,-4:

	S..............................................................
	...............................................................
	...............................................................
	...............................................................
	.................#.............................................
	....................TTTTTTTTTTT................................
	....................TTTTTTTTTTT................................
	....................TTTTTTTTTTT................................
	....................TTTTTTTTTTT................................
	....................TTTTTTTTTTT..#.............................
	....................TTTTTTTTTTT................................
	...............................................................
	...............................................................
	...............................................................
	...............................................................
	................................................#..............
	...............................................................
	...............................................................
	...............................................................
	...............................................................
	...............................................................
	...............................................................
	..............................................................#

The probe appears to pass through the target area, but is never within it after
any step. Instead, it continues down and to the right - only the first few
steps are shown.

If you're going to fire a highly scientific probe out of a super cool probe
launcher, you might as well do it with style. How high can you make the probe
go while still reaching the target area?

In the above example, using an initial velocity of 6,9 is the best you can do,
causing the probe to reach a maximum y position of 45. (Any higher initial y
velocity causes the probe to overshoot the target area entirely.)

Find the initial velocity that causes the probe to reach the highest y position
and still eventually be within the target area after any step. What is the
highest y position it reaches on this trajectory?
*/
func Part1(r io.Reader) (ans int, err error) {
	t, err := getTarget(r)
	if err != nil {
		return 0, err
	}
	height, _ := search(t)
	return height, nil
}

/*
Part2 Prompt

--- Part Two ---
Maybe a fancy trick shot isn't the best idea; after all, you only have one
probe, so you had better not miss.

To get the best idea of what your options are for launching the probe, you need
to find every initial velocity that causes the probe to eventually be within
the target area after any step.

In the above example, there are 112 different initial velocity values that meet
these criteria:

	23,-10  25,-9   27,-5   29,-6   22,-6   21,-7   9,0     27,-7   24,-5
	25,-7   26,-6   25,-5   6,8     11,-2   20,-5   29,-10  6,3     28,-7
	8,0     30,-6   29,-8   20,-10  6,7     6,4     6,1     14,-4   21,-6
	26,-10  7,-1    7,7     8,-1    21,-9   6,2     20,-7   30,-10  14,-3
	20,-8   13,-2   7,3     28,-8   29,-9   15,-3   22,-5   26,-8   25,-8
	25,-6   15,-4   9,-2    15,-2   12,-2   28,-9   12,-3   24,-6   23,-7
	25,-10  7,8     11,-3   26,-7   7,1     23,-9   6,0     22,-10  27,-6
	8,1     22,-8   13,-4   7,6     28,-6   11,-4   12,-4   26,-9   7,4
	24,-10  23,-8   30,-8   7,0     9,-1    10,-1   26,-5   22,-9   6,5
	7,5     23,-6   28,-10  10,-2   11,-1   20,-9   14,-2   29,-7   13,-3
	23,-5   24,-8   27,-9   30,-7   28,-5   21,-10  7,9     6,6     21,-5
	27,-10  7,2     30,-9   21,-8   22,-7   24,-9   20,-6   6,9     29,-5
	8,-2    27,-8   30,-5   24,-7

How many distinct initial velocity values cause the probe to be within the
target area after any step?
*/
func Part2(r io.Reader) (ans int, err error) {
	t, err := getTarget(r)
	if err != nil {
		return 0, err
	}
	_, hits := search(t)
	return hits, nil
}

// target area: x=20..30, y=-10..-5
var re = regexp.MustCompile(`target area: x=(.+)\.\.(.+), y=(.+)\.\.(.+)`)

type target struct {
	x1, x2, y1, y2 int
}

func (t target) contains(x, y int) bool {
	return x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2
}

func getTarget(r io.Reader) (target, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return target{}, err
	}
	m := re.FindStringSubmatch(string(b))
	if m == nil {
		return target{}, fmt.Errorf("bad input")
	}
	x1, x2, y1, y2 := num(m[1]), num(m[2]), num(m[3]), num(m[4])
	if x2 < 0 {
		x1, x2 = -x2, -x1
	}
	t := target{x1, x2, y1, y2}
	return t, nil
}

// We try all combinations of initial velocities that could possibly hit the target by trying all combinations of
// viable vx and vy. We know the target extends to the right because during parsing if it's to the left the
// coordinates are flipped. So we only try 0 to x2 for the vx. We bound it to x2 because if vx > x2, the first step
// will be beyond the target. Similar for vy, we use a range of y1 (which is negative, so it would basically be shooting
// at the target directly), an no lower because it would shoot past it on the first step, all the way to -y1. We know
// we can bound our search to this because of the symmetry of the changes in vy, the path will pass through y=0 with
// the opposite of the initial vy, so any higher of an initial vy will come back down just as fast and at the next step
// after passing y=0 it will be past the target.
func search(t target) (maxHeight, numHits int) {
	for vy := t.y1; vy <= -t.y1; vy++ {
		for vx := 0; vx <= t.x2; vx++ {
			h, hit := simulate(t, vx, vy)
			if hit {
				numHits++
				if h > maxHeight {
					maxHeight = h
				}
			}
		}
	}
	return maxHeight, numHits
}

func simulate(t target, vx, vy int) (height int, hit bool) {
	x, y := 0, 0
	for y >= t.y1 {
		if y > height {
			height = y
		}
		if t.contains(x, y) {
			return height, true
		}
		dvx := -1
		if vx == 0 {
			dvx = 0
		}
		x, y, vx, vy = x+vx, y+vy, vx+dvx, vy-1
	}
	return 0, false
}

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("bad input")
	}
	return n
}
