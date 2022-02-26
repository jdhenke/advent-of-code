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

const part2Signature = ``

func Part2(r io.Reader) (answer int, err error) {
	grid, err := solve(r)
	if err != nil {
		return 0, err
	}
	t := Tile{
		Data: strings.Split(strings.TrimSpace(part2Signature), "\n"),
	}
	t.ForAllOrientations(func(t Tile) {
		grid.ForAllLocations(func(i, j int) {
			if grid.MatchesSignature(i, j, t.Data) {
				answer++
			}
		})
	})
	return answer, nil
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
			if sig[i][j] != '#' {
				continue
			}
			if g.at(i+si, j+sj) != '#' {
				return false
			}
		}
	}
	return true
}

func (g Grid) at(i int, j int) byte {
	m := len(g[0][0].Data)
	ti := i / m
	n := len(g[0][0].Data[0])
	tj := j / n
	if ti >= len(g) || tj > len(g[0]) {
		return '\x00'
	}
	return g[ti][tj].Data[i%m][j%n]
}

type Tile struct {
	ID   int
	Data []string
}

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

	// create index of BorderAndSides to tiles that match that
	m := make(mapping)
	for _, t := range tiles {
		t.ForAllOrientations(func(t Tile) {
			t.ForAllBorders(func(bs BorderAndSide) {
				m[bs] = append(m[bs], t)
			})
		})
	}

	var corners []Tile
	for bs, ts := range m {
		// If the only matching tile to its left would be itself (which will always exist)
		if bs.Side == SideTop && len(ts) == 1 && len(m[ts[0].Match(SideLeft)]) == 1 {
			corners = append(corners, ts[0])
		}
	}
	// all four corner tiles in original or flipped could be top left ==> 8 possible tiles
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

type mapping map[BorderAndSide][]Tile

var re = regexp.MustCompile(`Tile (\d+):`)

func parseTileNum(s string) (int, error) {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return 0, fmt.Errorf("bad tile header: %v", s)
	}
	return strconv.Atoi(match[1])
}
