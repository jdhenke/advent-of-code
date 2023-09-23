package day24

import (
	"io"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 24: Lobby Layout ---
Your raft makes it to the tropical island; it turns out that the small crab was
an excellent navigator. You make your way to the resort.

As you enter the lobby, you discover a small problem: the floor is being
renovated. You can't even reach the check-in desk until they've finished
installing the new tile floor.

The tiles are all hexagonal; they need to be arranged in a hex grid with a very
specific color pattern. Not in the mood to wait, you offer to help figure out
the pattern.

The tiles are all white on one side and black on the other. They start with the
white side facing up. The lobby is large enough to fit whatever pattern might
need to appear there.

A member of the renovation crew gives you a list of the tiles that need to be
flipped over (your puzzle input). Each line in the list identifies a single
tile that needs to be flipped by giving a series of steps starting from a
reference tile in the very center of the room. (Every line starts from the same
reference tile.)

Because the tiles are hexagonal, every tile has six neighbors: east, southeast,
southwest, west, northwest, and northeast. These directions are given in your
list, respectively, as e, se, sw, w, nw, and ne. A tile is identified by a
series of these directions with no delimiters; for example, esenee identifies
the tile you land on if you start at the reference tile and then move one tile
east, one tile southeast, one tile northeast, and one tile east.

Each time a tile is identified, it flips from white to black or from black to
white. Tiles might be flipped more than once. For example, a line like esew
flips a tile immediately adjacent to the reference tile, and a line like
nwwswee flips the reference tile itself.

Here is a larger example:

	sesenwnenenewseeswwswswwnenewsewsw
	neeenesenwnwwswnenewnwwsewnenwseswesw
	seswneswswsenwwnwse
	nwnwneseeswswnenewneswwnewseswneseene
	swweswneswnenwsewnwneneseenw
	eesenwseswswnenwswnwnwsewwnwsene
	sewnenenenesenwsewnenwwwse
	wenwwweseeeweswwwnwwe
	wsweesenenewnwwnwsenewsenwwsesesenwne
	neeswseenwwswnwswswnw
	nenwswwsewswnenenewsenwsenwnesesenew
	enewnwewneswsewnwswenweswnenwsenwsw
	sweneswneswneneenwnewenewwneswswnese
	swwesenesewenwneswnwwneseswwne
	enesenwswwswneneswsenwnewswseenwsese
	wnwnesenesenenwwnenwsewesewsesesew
	nenewswnwewswnenesenwnesewesw
	eneswnwswnwsenenwnwnwwseeswneewsenese
	neswnwewnwnwseenwseesewsenwsweewe
	wseweeenwnesenwwwswnew

In the above example, 10 tiles are flipped once (to black), and 5 more are
flipped twice (to black, then back to white). After all of these instructions
have been followed, a total of 10 tiles are black.

Go through the renovation crew's list and determine which tiles they need to
flip. After all of the instructions have been followed, how many tiles are left
with the black side up?
*/
func Part1(r io.Reader) (answer int, err error) {
	locs, err := day24(r)
	if err != nil {
		return 0, err
	}
	return len(locs), nil
}

/*
Part2 Prompt

--- Part Two ---
The tile floor in the lobby is meant to be a living art exhibit. Every day, the
tiles are all flipped according to the following rules:

- Any black tile with zero or more than 2 black tiles immediately adjacent to
it is flipped to white.
- Any white tile with exactly 2 black tiles immediately adjacent to it is
flipped to black.

Here, tiles immediately adjacent means the six tiles directly touching the tile
in question.

The rules are applied simultaneously to every tile; put another way, it is
first determined which tiles need to be flipped, then they are all flipped at
the same time.

In the above example, the number of black tiles that are facing up after the
given number of days has passed is as follows:

	Day 1: 15
	Day 2: 12
	Day 3: 25
	Day 4: 14
	Day 5: 23
	Day 6: 28
	Day 7: 41
	Day 8: 37
	Day 9: 49
	Day 10: 37

	Day 20: 132
	Day 30: 259
	Day 40: 406
	Day 50: 566
	Day 60: 788
	Day 70: 1106
	Day 80: 1373
	Day 90: 1844
	Day 100: 2208

After executing this process a total of 100 times, there would be 2208 black
tiles facing up.

How many tiles will be black after 100 days?
*/
func Part2(r io.Reader) (answer int, err error) {
	// initialize using part 1
	locs, err := day24(r)
	if err != nil {
		return 0, err
	}

	// step through 100 days of rules from part 2
	for i := 0; i < 100; i++ {
		// get bounding box of known black tiles at this point
		var minX, maxX, minY, maxY int
		for l := range locs {
			updateBounds(&minX, &maxX, l.x)
			updateBounds(&minY, &maxY, l.y)
		}
		toFlip := make(map[loc]bool)

		// check all tiles that may be adjacent to any black tile
		for x := minX - 2; x <= maxX+2; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				n := neighbors(locs, x, y)
				l := loc{x, y}
				if _, ok := locs[l]; ok && (n == 0 || n > 2) {
					toFlip[l] = false
				} else if !ok && (n == 2) {
					toFlip[l] = true
				}
			}
		}

		// flip them after calculating them at all at once
		for l, v := range toFlip {
			if v {
				locs[l] = struct{}{}
			} else {
				delete(locs, l)
			}
		}
	}
	return len(locs), nil
}

type loc struct {
	x, y int
}

type change struct {
	dx, dy int
}

var instructions = map[string]change{
	"e":  {2, 0},
	"se": {1, -1},
	"sw": {-1, -1},
	"w":  {-2, 0},
	"nw": {-1, 1},
	"ne": {1, 1},
}

func updateBounds(min *int, max *int, x int) {
	if x < *min {
		*min = x
	}
	if x > *max {
		*max = x
	}
}

func neighbors(locs map[loc]struct{}, x, y int) int {
	n := 0
	for _, delta := range instructions {
		if _, ok := locs[loc{x + delta.dx, y + delta.dy}]; ok {
			n++
		}
	}
	return n
}

func day24(r io.Reader) (locs map[loc]struct{}, err error) {
	locs = make(map[loc]struct{})
	if err := input.ForEachLine(r, func(line string) error {
		x, y := 0, 0
		forEachInstruction(line, func(delta change) {
			x += delta.dx
			y += delta.dy
		})
		l := loc{x, y}
		if _, ok := locs[l]; ok {
			delete(locs, l)
		} else {
			locs[l] = struct{}{}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return locs, nil
}

func forEachInstruction(line string, f func(delta change)) {
	i := 0
	for i < len(line) {
		for prefix, delta := range instructions {
			if strings.HasPrefix(line[i:], prefix) {
				f(delta)
				i += len(prefix)
			}
		}
	}
}
