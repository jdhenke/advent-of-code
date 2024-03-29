package day20

import (
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 20: Jurassic Jigsaw ---
The high-speed train leaves the forest and quickly carries you south. You can
even see a desert in the distance! Since you have some spare time, you might as
well see if there was anything interesting in the image the Mythical
Information Bureau satellite captured.

After decoding the satellite messages, you discover that the data actually
contains many small images created by the satellite's camera array. The camera
array consists of many cameras; rather than produce a single square image, they
produce many smaller square image tiles that need to be reassembled back into a
single image.

Each camera in the camera array returns a single monochrome image tile with a
random unique ID number. The tiles (your puzzle input) arrived in a random
order.

Worse yet, the camera array appears to be malfunctioning: each image tile has
been rotated and flipped to a random orientation. Your first task is to
reassemble the original image by orienting the tiles so they fit together.

To show how the tiles should be reassembled, each tile's image data includes a
border that should line up exactly with its adjacent tiles. All tiles have this
border, and the border lines up exactly when the tiles are both oriented
correctly. Tiles at the edge of the image also have this border, but the
outermost edges won't line up with any other tiles.

For example, suppose you have the following nine tiles:

	Tile 2311:
	..##.#..#.
	##..#.....
	#...##..#.
	####.#...#
	##.##.###.
	##...#.###
	.#.#.#..##
	..#....#..
	###...#.#.
	..###..###

	Tile 1951:
	#.##...##.
	#.####...#
	.....#..##
	#...######
	.##.#....#
	.###.#####
	###.##.##.
	.###....#.
	..#.#..#.#
	#...##.#..

	Tile 1171:
	####...##.
	#..##.#..#
	##.#..#.#.
	.###.####.
	..###.####
	.##....##.
	.#...####.
	#.##.####.
	####..#...
	.....##...

	Tile 1427:
	###.##.#..
	.#..#.##..
	.#.##.#..#
	#.#.#.##.#
	....#...##
	...##..##.
	...#.#####
	.#.####.#.
	..#..###.#
	..##.#..#.

	Tile 1489:
	##.#.#....
	..##...#..
	.##..##...
	..#...#...
	#####...#.
	#..#.#.#.#
	...#.#.#..
	##.#...##.
	..##.##.##
	###.##.#..

	Tile 2473:
	#....####.
	#..#.##...
	#.##..#...
	######.#.#
	.#...#.#.#
	.#########
	.###.#..#.
	########.#
	##...##.#.
	..###.#.#.

	Tile 2971:
	..#.#....#
	#...###...
	#.#.###...
	##.##..#..
	.#####..##
	.#..####.#
	#..#.#..#.
	..####.###
	..#.#.###.
	...#.#.#.#

	Tile 2729:
	...#.#.#.#
	####.#....
	..#.#.....
	....#..#.#
	.##..##.#.
	.#.####...
	####.#.#..
	##.####...
	##..#.##..
	#.##...##.

	Tile 3079:
	#.#.#####.
	.#..######
	..#.......
	######....
	####.#..#.
	.#...#.##.
	#.#####.##
	..#.###...
	..#.......
	..#.###...

By rotating, flipping, and rearranging them, you can find a square arrangement
that causes all adjacent borders to line up:

	#...##.#.. ..###..### #.#.#####.
	..#.#..#.# ###...#.#. .#..######
	.###....#. ..#....#.. ..#.......
	###.##.##. .#.#.#..## ######....
	.###.##### ##...#.### ####.#..#.
	.##.#....# ##.##.###. .#...#.##.
	#...###### ####.#...# #.#####.##
	.....#..## #...##..#. ..#.###...
	#.####...# ##..#..... ..#.......
	#.##...##. ..##.#..#. ..#.###...

	#.##...##. ..##.#..#. ..#.###...
	##..#.##.. ..#..###.# ##.##....#
	##.####... .#.####.#. ..#.###..#
	####.#.#.. ...#.##### ###.#..###
	.#.####... ...##..##. .######.##
	.##..##.#. ....#...## #.#.#.#...
	....#..#.# #.#.#.##.# #.###.###.
	..#.#..... .#.##.#..# #.###.##..
	####.#.... .#..#.##.. .######...
	...#.#.#.# ###.##.#.. .##...####

	...#.#.#.# ###.##.#.. .##...####
	..#.#.###. ..##.##.## #..#.##..#
	..####.### ##.#...##. .#.#..#.##
	#..#.#..#. ...#.#.#.. .####.###.
	.#..####.# #..#.#.#.# ####.###..
	.#####..## #####...#. .##....##.
	##.##..#.. ..#...#... .####...#.
	#.#.###... .##..##... .####.##.#
	#...###... ..##...#.. ...#..####
	..#.#....# ##.#.#.... ...##.....

For reference, the IDs of the above tiles are:

	1951    2311    3079
	2729    1427    2473
	2971    1489    1171

To check that you've assembled the image correctly, multiply the IDs of the
four corner tiles together. If you do this with the assembled tiles from the
example above, you get 1951 * 3079 * 2971 * 1171 = 20899048083289.

Assemble the tiles into an image. What do you get if you multiply together the
IDs of the four corner tiles?
*/
func Part1(r io.Reader) (answer int, err error) {
	grid, err := solve(r)
	if err != nil {
		return 0, err
	}
	answer = 1
	m, n := len(grid), len(grid[0])
	answer *= grid[0][0].ID
	answer *= grid[0][n-1].ID
	answer *= grid[m-1][0].ID
	answer *= grid[m-1][n-1].ID
	return answer, nil
}

const part2Signature = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

/*
Part2 Prompt

--- Part Two ---
Now, you're ready to check the image for sea monsters.

The borders of each tile are not part of the actual image; start by removing
them.

In the example above, the tiles become:

	.#.#..#. ##...#.# #..#####
	###....# .#....#. .#......
	##.##.## #.#.#..# #####...
	###.#### #...#.## ###.#..#
	##.#.... #.##.### #...#.##
	...##### ###.#... .#####.#
	....#..# ...##..# .#.###..
	.####... #..#.... .#......

	#..#.##. .#..###. #.##....
	#.####.. #.####.# .#.###..
	###.#.#. ..#.#### ##.#..##
	#.####.. ..##..## ######.#
	##..##.# ...#...# .#.#.#..
	...#..#. .#.#.##. .###.###
	.#.#.... #.##.#.. .###.##.
	###.#... #..#.##. ######..

	.#.#.### .##.##.# ..#.##..
	.####.## #.#...## #.#..#.#
	..#.#..# ..#.#.#. ####.###
	#..####. ..#.#.#. ###.###.
	#####..# ####...# ##....##
	#.##..#. .#...#.. ####...#
	.#.###.. ##..##.. ####.##.
	...###.. .##...#. ..#..###

Remove the gaps to form the actual image:

	.#.#..#.##...#.##..#####
	###....#.#....#..#......
	##.##.###.#.#..######...
	###.#####...#.#####.#..#
	##.#....#.##.####...#.##
	...########.#....#####.#
	....#..#...##..#.#.###..
	.####...#..#.....#......
	#..#.##..#..###.#.##....
	#.####..#.####.#.#.###..
	###.#.#...#.######.#..##
	#.####....##..########.#
	##..##.#...#...#.#.#.#..
	...#..#..#.#.##..###.###
	.#.#....#.##.#...###.##.
	###.#...#..#.##.######..
	.#.#.###.##.##.#..#.##..
	.####.###.#...###.#..#.#
	..#.#..#..#.#.#.####.###
	#..####...#.#.#.###.###.
	#####..#####...###....##
	#.##..#..#...#..####...#
	.#.###..##..##..####.##.
	...###...##...#...#..###

Now, you're ready to search for sea monsters! Because your image is monochrome,
a sea monster will look like this:

	                  #
	#    ##    ##    ###
	 #  #  #  #  #  #

When looking for this pattern in the image, the spaces can be anything; only
the # need to match. Also, you might need to rotate or flip your image before
it's oriented correctly to find sea monsters. In the above image, after
flipping and rotating it to the appropriate orientation, there are two sea
monsters (marked with O):

	.####...#####..#...###..
	#####..#..#.#.####..#.#.
	.#.#...#.###...#.##.O#..
	#.O.##.OO#.#.OO.##.OOO##
	..#O.#O#.O##O..O.#O##.##
	...#.#..##.##...#..#..##
	#.##.#..#.#..#..##.#.#..
	.###.##.....#...###.#...
	#.####.#.#....##.#..#.#.
	##...#..#....#..#...####
	..#.##...###..#.#####..#
	....#.##.#.#####....#...
	..##.##.###.....#.##..#.
	#...#...###..####....##.
	.#.##...#.##.#.#.###...#
	#.###.#..####...##..#...
	#.###...#.##...#.##O###.
	.O##.#OO.###OO##..OOO##.
	..O#.O..O..O.#O##O##.###
	#.#..##.########..#..##.
	#.#####..#.#...##..#....
	#....##..#.#########..##
	#...#.....#..##...###.##
	#..###....##.#...##.##.#

Determine how rough the waters are in the sea monsters' habitat by counting the
number of # that are not part of a sea monster. In the above example, the
habitat's water roughness is 273.

How many # are not part of a sea monster?
*/
func Part2(r io.Reader) (answer int, err error) {
	grid, err := solve(r)
	if err != nil {
		return 0, err
	}

	// Create a tile of the signature so Tile.ForAllOrientations can be reused to check all orientations of the
	// signature. This also avoids having to merge the existing tiles.
	t := Tile{
		Data: strings.Split(part2Signature, "\n"),
	}

	// Removes the borders which are not part of the image.
	grid.trim()

	// Assuming no overlapping signatures, calculate the number of '#' not in a signature by counting all '#' then
	// subtracting the number of signatures times the number of '#' in each signature.
	grid.ForAllLocations(func(i int, j int) {
		if grid.at(i, j) == '#' {
			answer++
		}
	})
	numMatches := 0
	t.ForAllOrientations(func(t Tile) {
		grid.ForAllLocations(func(i, j int) {
			if grid.MatchesSignature(i, j, t.Data) {
				numMatches++
			}
		})
	})
	answer -= numMatches * strings.Count(part2Signature, "#")
	return answer, nil
}

func solve(r io.Reader) (grid Grid, err error) {
	// parse tiles
	var tiles []Tile
	if err := input.ForEachBatch(r, func(batch []string) error {
		id, err := parseTileNum(batch[0])
		if err != nil {
			return err
		}
		tiles = append(tiles, Tile{
			ID:   id,
			Data: batch[1:],
		})
		return nil
	}); err != nil {
		return nil, err
	}

	// Create index of BorderAndSide to all orientations of all tiles. Note that this means every entry will end up
	// with at least one entry, the flipped version of the same tile.
	m := make(map[BorderAndSide][]Tile)
	for _, t := range tiles {
		t.ForAllOrientations(func(t Tile) {
			t.ForAllBorders(func(bs BorderAndSide) {
				m[bs] = append(m[bs], t)
			})
		})
	}

	// Identify an arbitrary, but deterministic, top left tile.
	var corners []Tile
	for bs, ts := range m {
		// If the tile corresponding to this top border only matches the flipped version of itself to the left, this is
		// a corner tile.
		if bs.Side == SideTop && len(ts) == 1 && len(m[ts[0].Match(SideLeft)]) == 1 {
			corners = append(corners, ts[0])
		}
	}
	// All four corner tiles in the original or flipped version could be top left in one of their rotations, so verify
	// all of these have in fact been found.
	if len(corners) != 8 {
		return nil, fmt.Errorf("could not identify corners: %d", len(corners))
	}
	sort.Slice(corners, func(i, j int) bool {
		if corners[i].ID != corners[j].ID {
			return corners[i].ID < corners[j].ID
		}
		return corners[i].Data[0] < corners[j].Data[0]
	})
	topLeft := corners[0]

	// Assemble the grid row by row, finding tiles that match to the right that have not been used yet, avoiding
	// including flipped or rotated versions of a tile that has already been used.
	seen := make(map[int]bool)
	grid = append(grid, []Tile{topLeft})
	seen[topLeft.ID] = true
	current := topLeft
	getUnseenMatch := func(bs BorderAndSide) (Tile, bool) {
		for _, match := range m[bs] {
			if seen[match.ID] {
				continue
			}
			seen[match.ID] = true
			return match, true
		}
		return Tile{}, false
	}
	for {
		var ok bool
		for current, ok = getUnseenMatch(current.Match(SideRight)); ok; current, ok = getUnseenMatch(current.Match(SideRight)) {
			grid[len(grid)-1] = append(grid[len(grid)-1], current)
		}
		current, ok = getUnseenMatch(grid[len(grid)-1][0].Match(SideBottom))
		if !ok {
			break
		}
		grid = append(grid, []Tile{current})
	}
	return grid, nil
}

var re = regexp.MustCompile(`Tile (\d+):`)

func parseTileNum(s string) (int, error) {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return 0, fmt.Errorf("bad tile header: %v", s)
	}
	return strconv.Atoi(match[1])
}

type Grid [][]Tile

func (g Grid) ForAllLocations(f func(i int, j int)) {
	for ti, row := range g {
		for tj, t := range row {
			for i := range g[ti][tj].Data {
				for j := range g[ti][tj].Data[i] {
					f(ti*len(t.Data)+i, tj*len(t.Data[0])+j)
				}
			}
		}
	}
}

func (g Grid) MatchesSignature(i int, j int, sig []string) bool {
	for si := 0; si < len(sig); si++ {
		for sj := 0; sj < len(sig[0]); sj++ {
			if sig[si][sj] != '#' {
				continue
			}
			if g.at(i+si, j+sj) != '#' {
				return false
			}
		}
	}
	return true
}

// assumes all tiles are the same size, returns a null byte if out of bounds of the grid.
func (g Grid) at(i int, j int) byte {
	m := len(g[0][0].Data)
	ti := i / m
	n := len(g[0][0].Data[0])
	tj := j / n
	if ti >= len(g) || tj >= len(g[0]) {
		return '\x00'
	}
	return g[ti][tj].Data[i%m][j%n]
}

func (g Grid) trim() {
	for gi := range g {
		for gj := range g[gi] {
			var newData []string
			for i := 1; i < len(g[gi][gj].Data)-1; i++ {
				newData = append(newData, g[gi][gj].Data[i][1:len(g[gi][gj].Data)-1])
			}
			g[gi][gj].Data = newData
		}
	}
}

type Tile struct {
	ID   int
	Data []string
}

// Match returns a BorderAndSide that would match this Tile's given Side.
//
// For example, calling Match(SideRight) for this Tile:
//
//	A B C
//	D E F
//	G H I
//
// Would return SideLeft for "CFI".
func (t Tile) Match(side Side) BorderAndSide {
	matchingSide := map[Side]Side{
		SideTop:    SideBottom,
		SideLeft:   SideRight,
		SideBottom: SideTop,
		SideRight:  SideLeft,
	}[side]
	rowFunc := func(i int) func() string {
		return func() string {
			return t.row(i)
		}
	}
	colFunc := func(j int) func() string {
		return func() string {
			return t.col(j)
		}
	}
	border := map[Side]func() string{
		SideTop:    rowFunc(0),
		SideBottom: rowFunc(len(t.Data) - 1),
		SideLeft:   colFunc(0),
		SideRight:  colFunc(len(t.Data[0]) - 1),
	}[side]()
	return BorderAndSide{
		Side:   matchingSide,
		Border: border,
	}
}

func (t Tile) row(i int) string {
	return t.Data[i]
}

func (t Tile) col(j int) string {
	var buf strings.Builder
	for i := 0; i < len(t.Data); i++ {
		buf.WriteByte(t.Data[i][j])
	}
	return buf.String()
}

// ForAllOrientations calls f for each of the eight different flipped and rotated versions of t.
func (t Tile) ForAllOrientations(f func(t Tile)) {
	for _, temp := range []Tile{t, t.flip()} {
		for i := 0; i < 4; i++ {
			f(temp)
			temp = temp.rotate()
		}
	}
}

// flip over the i = j line
func (t Tile) flip() Tile {
	out := Tile{
		ID: t.ID,
	}
	m, n := len(t.Data), len(t.Data[0])
	for j := 0; j < n; j++ {
		row := &strings.Builder{}
		for i := 0; i < m; i++ {
			row.WriteByte(t.Data[i][j])
		}
		out.Data = append(out.Data, row.String())
	}
	return out
}

// rotate 90 degrees clockwise
func (t Tile) rotate() Tile {
	out := Tile{
		ID: t.ID,
	}
	m, n := len(t.Data), len(t.Data[0])
	for j := 0; j < n; j++ {
		row := &strings.Builder{}
		for i := m - 1; i >= 0; i-- {
			row.WriteByte(t.Data[i][j])
		}
		out.Data = append(out.Data, row.String())
	}
	return out
}

func (t Tile) ForAllBorders(f func(bs BorderAndSide)) {
	for _, bs := range []BorderAndSide{
		{
			Side:   SideTop,
			Border: t.row(0),
		},
		{
			Side:   SideRight,
			Border: t.col(len(t.Data[0]) - 1),
		},
		{
			Side:   SideBottom,
			Border: t.row(len(t.Data) - 1),
		},
		{
			Side:   SideLeft,
			Border: t.col(0),
		},
	} {
		f(bs)
	}
}

type Side int

const (
	SideTop    Side = iota // text is left to right
	SideRight              // text is top down
	SideBottom             // text is left to right
	SideLeft               // text is top down
)

type BorderAndSide struct {
	Side   Side
	Border string
}
