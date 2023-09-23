package day23

import (
	"container/heap"
	"fmt"
	"io"
	"regexp"
	"sort"
)

// Board represents the state of the board.
//
// The different fields are ints representing the location of that type of amphipod on the board e.g. A1 is the first A,
// B2 is the second B, etc...
//
// The values of these fields correspond to the locations in the following way:
//
// In the "hallway" the positions are 0,1,2,3,4,5,6 as shown:
//
//	#############
//	#01.2.3.4.56#
//	###A#B#C#D### 0 (1s)
//	  #A#B#C#D#   1
//	  #A#B#C#D#   2
//	  #A#B#C#D#   3
//	  #########
//	   1 2 3 4    <-- (10s)
//
// In each of their "homes", the positions have a tens value of 1 for A's home, 2 for B's, etc... and a ones value of
// 0 right next to the hallway, 1 the next position down, etc... So for example, 23 would mean B's bottom most home
// position.
type Board struct {
	A1, A2, B1, B2, C1, C2, D1, D2 int
	A3, A4, B3, B4, C3, C4, D3, D4 int
}

/*
Part1 Prompt

--- Day 23: Amphipod ---
A group of amphipods notice your fancy submarine and flag you down. "With such
an impressive shell," one amphipod says, "surely you can help us with a
question that has stumped our best scientists."

They go on to explain that a group of timid, stubborn amphipods live in a
nearby burrow. Four types of amphipods live there: Amber (A), Bronze (B),
Copper (C), and Desert (D). They live in a burrow that consists of a hallway
and four side rooms. The side rooms are initially full of amphipods, and the
hallway is initially empty.

They give you a diagram of the situation (your puzzle input), including
locations of each amphipod (A, B, C, or D, each of which is occupying an
otherwise open space), walls (#), and open space (.).

For example:

	#############
	#...........#
	###B#C#B#D###
	  #A#D#C#A#
	  #########

The amphipods would like a method to organize every amphipod into side rooms so
that each side room contains one type of amphipod and the types are sorted A-D
going left to right, like this:

	#############
	#...........#
	###A#B#C#D###
	  #A#B#C#D#
	  #########

Amphipods can move up, down, left, or right so long as they are moving into an
unoccupied open space. Each type of amphipod requires a different amount of
energy to move one step: Amber amphipods require 1 energy per step, Bronze
amphipods require 10 energy, Copper amphipods require 100, and Desert ones
require 1000. The amphipods would like you to find a way to organize the
amphipods that requires the least total energy.

However, because they are timid and stubborn, the amphipods have some extra
rules:

- Amphipods will never stop on the space immediately outside any room. They can
move into that space so long as they immediately continue moving.
(Specifically, this refers to the four open spaces in the hallway that are
directly above an amphipod starting position.)
- Amphipods will never move from the hallway into a room unless that room is
their destination room and that room contains no amphipods which do not also
have that room as their own destination. If an amphipod's starting room is not
its destination room, it can stay in that room until it leaves the room. (For
example, an Amber amphipod will not move from the hallway into the right three
rooms, and will only move into the leftmost room if that room is empty or if it
only contains other Amber amphipods.)
- Once an amphipod stops moving in the hallway, it will stay in that spot until
it can move into a room. (That is, once any amphipod starts moving, any other
amphipods currently in the hallway are locked in place and will not move again
until they can move fully into a room.)

In the above example, the amphipods can be organized using a minimum of 12521
energy. One way to do this is shown below.

Starting configuration:

	#############
	#...........#
	###B#C#B#D###
	  #A#D#C#A#
	  #########

One Bronze amphipod moves into the hallway, taking 4 steps and using 40 energy:

	#############
	#...B.......#
	###B#C#.#D###
	  #A#D#C#A#
	  #########

The only Copper amphipod not in its side room moves there, taking 4 steps and
using 400 energy:

	#############
	#...B.......#
	###B#.#C#D###
	  #A#D#C#A#
	  #########

A Desert amphipod moves out of the way, taking 3 steps and using 3000 energy,
and then the Bronze amphipod takes its place, taking 3 steps and using 30
energy:

	#############
	#.....D.....#
	###B#.#C#D###
	  #A#B#C#A#
	  #########

The leftmost Bronze amphipod moves to its room using 40 energy:

	#############
	#.....D.....#
	###.#B#C#D###
	  #A#B#C#A#
	  #########

Both amphipods in the rightmost room move into the hallway, using 2003 energy
in total:

	#############
	#.....D.D.A.#
	###.#B#C#.###
	  #A#B#C#.#
	  #########

Both Desert amphipods move into the rightmost room using 7000 energy:

	#############
	#.........A.#
	###.#B#C#D###
	  #A#B#C#D#
	  #########

Finally, the last Amber amphipod moves into its room, using 8 energy:

	#############
	#...........#
	###A#B#C#D###
	  #A#B#C#D#
	  #########

What is the least energy required to organize the amphipods?
*/
func Part1(r io.Reader) (answer int, err error) {
	b, err := parseBoard(r)
	if err != nil {
		return 0, err
	}

	ans, ok := solve(b)
	if !ok {
		return 0, fmt.Errorf("could not solve board")
	}
	return ans, nil
}

/*
Part2 Prompt

--- Part Two ---
As you prepare to give the amphipods your solution, you notice that the diagram
they handed you was actually folded up. As you unfold it, you discover an extra
part of the diagram.

Between the first and second lines of text that contain amphipod starting
positions, insert the following lines:

	#D#C#B#A#
	#D#B#A#C#

So, the above example now becomes:

	#############
	#...........#
	###B#C#B#D###
	  #D#C#B#A#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

The amphipods still want to be organized into rooms similar to before:

	#############
	#...........#
	###A#B#C#D###
	  #A#B#C#D#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

In this updated example, the least energy required to organize these amphipods
is 44169:

	#############
	#...........#
	###B#C#B#D###
	  #D#C#B#A#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

	#############
	#..........D#
	###B#C#B#.###
	  #D#C#B#A#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

	#############
	#A.........D#
	###B#C#B#.###
	  #D#C#B#.#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

	#############
	#A........BD#
	###B#C#.#.###
	  #D#C#B#.#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

	#############
	#A......B.BD#
	###B#C#.#.###
	  #D#C#.#.#
	  #D#B#A#C#
	  #A#D#C#A#
	  #########

	#############
	#AA.....B.BD#
	###B#C#.#.###
	  #D#C#.#.#
	  #D#B#.#C#
	  #A#D#C#A#
	  #########

	#############
	#AA.....B.BD#
	###B#.#.#.###
	  #D#C#.#.#
	  #D#B#C#C#
	  #A#D#C#A#
	  #########

	#############
	#AA.....B.BD#
	###B#.#.#.###
	  #D#.#C#.#
	  #D#B#C#C#
	  #A#D#C#A#
	  #########

	#############
	#AA...B.B.BD#
	###B#.#.#.###
	  #D#.#C#.#
	  #D#.#C#C#
	  #A#D#C#A#
	  #########

	#############
	#AA.D.B.B.BD#
	###B#.#.#.###
	  #D#.#C#.#
	  #D#.#C#C#
	  #A#.#C#A#
	  #########

	#############
	#AA.D...B.BD#
	###B#.#.#.###
	  #D#.#C#.#
	  #D#.#C#C#
	  #A#B#C#A#
	  #########

	#############
	#AA.D.....BD#
	###B#.#.#.###
	  #D#.#C#.#
	  #D#B#C#C#
	  #A#B#C#A#
	  #########

	#############
	#AA.D......D#
	###B#.#.#.###
	  #D#B#C#.#
	  #D#B#C#C#
	  #A#B#C#A#
	  #########

	#############
	#AA.D......D#
	###B#.#C#.###
	  #D#B#C#.#
	  #D#B#C#.#
	  #A#B#C#A#
	  #########

	#############
	#AA.D.....AD#
	###B#.#C#.###
	  #D#B#C#.#
	  #D#B#C#.#
	  #A#B#C#.#
	  #########

	#############
	#AA.......AD#
	###B#.#C#.###
	  #D#B#C#.#
	  #D#B#C#.#
	  #A#B#C#D#
	  #########

	#############
	#AA.......AD#
	###.#B#C#.###
	  #D#B#C#.#
	  #D#B#C#.#
	  #A#B#C#D#
	  #########

	#############
	#AA.......AD#
	###.#B#C#.###
	  #.#B#C#.#
	  #D#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#AA.D.....AD#
	###.#B#C#.###
	  #.#B#C#.#
	  #.#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#A..D.....AD#
	###.#B#C#.###
	  #.#B#C#.#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#...D.....AD#
	###.#B#C#.###
	  #A#B#C#.#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#.........AD#
	###.#B#C#.###
	  #A#B#C#D#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#..........D#
	###A#B#C#.###
	  #A#B#C#D#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

	#############
	#...........#
	###A#B#C#D###
	  #A#B#C#D#
	  #A#B#C#D#
	  #A#B#C#D#
	  #########

Using the initial configuration from the full diagram, what is the least energy
required to organize the amphipods?
*/
func Part2(r io.Reader) (answer int, err error) {
	b, err := parseBoard(r)
	if err != nil {
		return 0, err
	}
	b = part2Shift(b)
	ans, ok := solve(b)
	if !ok {
		return 0, fmt.Errorf("could not solve board")
	}
	return ans, nil
}

func part2Shift(b Board) Board {
	// shift all X1's down to X3's
	for _, p := range []*int{
		&b.A1, &b.A2,
		&b.B1, &b.B2,
		&b.C1, &b.C2,
		&b.D1, &b.D2,
	} {
		if *p%10 == 1 {
			*p += 2
		}
	}
	b.D3 = 11
	b.D4 = 12
	b.C3 = 21
	b.B3 = 22
	b.B4 = 31
	b.A3 = 32
	b.A4 = 41
	b.C4 = 42
	return b
}

var re = regexp.MustCompile(`#############
#...........#
###(.)#(.)#(.)#(.)###
  #(.)#(.)#(.)#(.)#
  #########`)

// Returns a board parsed as in part 1 but with the size of part 2, filling in the bottom two rows as already being
// solved so the same logic can be used to solve both parts.
func parseBoard(r io.Reader) (Board, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return Board{}, err
	}
	parts := re.FindStringSubmatch(string(data))
	reverse := map[int]string{
		10: parts[1],
		20: parts[2],
		30: parts[3],
		40: parts[4],
		11: parts[5],
		21: parts[6],
		31: parts[7],
		41: parts[8],
	}
	lookup := make(map[string][]int)
	for loc, pieceStr := range reverse {
		lookup[pieceStr] = append(lookup[pieceStr], loc)
		sort.Ints(lookup[pieceStr])
	}
	b := Board{
		A1: lookup["A"][0],
		A2: lookup["A"][1],
		B1: lookup["B"][0],
		B2: lookup["B"][1],
		C1: lookup["C"][0],
		C2: lookup["C"][1],
		D1: lookup["D"][0],
		D2: lookup["D"][1],
	}
	b.D3 = 42
	b.D4 = 43
	b.C3 = 32
	b.B3 = 22
	b.B4 = 23
	b.A3 = 12
	b.A4 = 13
	b.C4 = 33
	return b, nil
}

type entry struct {
	board Board
	cost  int
}

type queue []entry

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].cost < q[j].cost
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(entry))
}

func (q *queue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}

var covered = make(map[Board]bool)

// dijkstra over all board states by cost to get there and return the first, therefore cheapest, solved board path cost.
func solve(b Board) (ans int, ok bool) {
	q := queue{entry{
		board: b,
		cost:  0,
	}}
	for len(q) > 0 {
		e := heap.Pop(&q).(entry)
		if covered[e.board] {
			continue
		}
		covered[e.board] = true
		if solved(e.board) {
			return e.cost, true
		}
		for nextBoard, moveCost := range getAllNext(e.board) {
			heap.Push(&q, entry{
				cost:  e.cost + moveCost,
				board: nextBoard,
			})
		}
	}
	return 0, false
}

func solved(b Board) bool {
	return inHome(1, b.A1, b.A2, b.A3, b.A4) &&
		inHome(2, b.B1, b.B2, b.B3, b.B4) &&
		inHome(3, b.C1, b.C2, b.C3, b.C4) &&
		inHome(4, b.D1, b.D2, b.D3, b.D4)
}

func inHome(h int, x1 int, x2 int, x3 int, x4 int) bool {
	return x1/10 == h && x2/10 == h && x3/10 == h && x4/10 == h
}

func getAllNext(b Board) map[Board]int {
	out := make(map[Board]int)
	for _, x := range []struct {
		loc         *int
		costPerStep int
	}{
		{&b.A1, 1},
		{&b.A2, 1},
		{&b.A3, 1},
		{&b.A4, 1},
		{&b.B1, 10},
		{&b.B2, 10},
		{&b.B3, 10},
		{&b.B4, 10},
		{&b.C1, 100},
		{&b.C2, 100},
		{&b.C3, 100},
		{&b.C4, 100},
		{&b.D1, 1000},
		{&b.D2, 1000},
		{&b.D3, 1000},
		{&b.D4, 1000},
	} {
		for loc, steps := range moves(b, *x.loc) {
			old := *x.loc
			*x.loc = loc
			out[b] = steps * x.costPerStep
			*x.loc = old
		}
	}
	return out
}

func moves(b Board, pos int) map[int]int {
	out := make(map[int]int)
	locs := map[int]int{
		b.A1: 1,
		b.A2: 1,
		b.A3: 1,
		b.A4: 1,
		b.B1: 2,
		b.B2: 2,
		b.B3: 2,
		b.B4: 2,
		b.C1: 3,
		b.C2: 3,
		b.C3: 3,
		b.C4: 3,
		b.D1: 4,
		b.D2: 4,
		b.D3: 4,
		b.D4: 4,
	}
	piece := locs[pos]
	try := func(start, end int) {
		steps, n := getPath(start, end)
		for _, loc := range steps {
			if locs[loc] != 0 {
				return
			}
		}
		out[end] = n
	}
	if pos < 10 {
		end := (piece * 10) + 3
		for locs[end] == piece {
			end--
		}
		try(pos, end)
	} else {
		finalSpot := true
		if pos/10 == piece {
			for x := (piece * 10) + 3; x >= pos; x-- {
				if locs[x] != piece {
					finalSpot = false
					break
				}
			}
		} else {
			finalSpot = false
		}
		if finalSpot {
			return nil
		}
		for dest := 0; dest <= 6; dest++ {
			try(pos, dest)
		}
	}
	return out
}

func getPath(a, b int) (path []int, n int) {
	if a == b {
		return nil, 0
	}
	if a > b {
		sp, n := getPath(b, a)
		for i := len(sp) - 2; i >= 0; i-- { // ignore a at end
			path = append(path, sp[i])
		}
		path = append(path, b) // add b that the reverse wouldn't
		return path, n
	}
	add := func(next int, cost int) (path []int, n int) {
		sp, sn := getPath(next, b)
		path = append(path, next)
		path = append(path, sp...)
		return path, cost + sn
	}
	if a == 0 {
		return add(1, 1)
	}
	switch a {
	case 0:
		return add(1, 1)
	case 1:
		switch b / 10 {
		case 1:
			return add(10, 2)
		}
		return add(2, 2)
	case 2:
		switch b / 10 {
		case 1:
			return add(10, 2)
		case 2:
			return add(20, 2)
		}
		return add(3, 2)
	case 3:
		switch b / 10 {
		case 1:
			return add(2, 2)
		case 2:
			return add(20, 2)
		case 3:
			return add(30, 2)
		}
		return add(4, 2)
	case 4:
		switch b / 10 {
		case 4:
			return add(40, 2)
		case 3:
			return add(30, 2)
		}
		return add(3, 2)
	case 5:
		switch b / 10 {
		case 4:
			return add(40, 2)
		}
		return add(4, 2)
	case 6:
		return add(5, 1)
	default: // in a home trying to get farther in
		return add(a+1, 1)
	}
}
