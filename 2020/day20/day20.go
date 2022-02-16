package day20

import (
	"fmt"
	"io"
	"regexp"
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
	return day20(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day20(r)
}

type Tile struct {
	ID   int
	Data []string
}

func (t Tile) Borders() []string {
	var out []string
	out = append(out, t.Data[0])
	out = append(out, t.Data[len(t.Data)-1])
	left, right := &strings.Builder{}, &strings.Builder{}
	for _, line := range t.Data {
		left.WriteString(line[:1])
		right.WriteString(line[len(line)-1:])
	}
	out = append(out, left.String())
	out = append(out, right.String())
	return out
}

func day20(r io.Reader) (answer int, err error) {
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
		return 0, err
	}
	// border matches
	borders := make(map[string][]int) // all possible border orientations to ==>  index of tiles that have that border
	for i, tile := range tiles {
		for _, border := range tile.Borders() {
			for _, s := range []string{
				border,
				reverse(border),
			} {
				borders[s] = append(borders[s], i)
			}
		}
	}

	// show how many matching borders there are
	counts := make(map[int]int) // tiles index ==> num pairs
	for _, is := range borders {
		if len(is) == 1 {
			continue // unmatched border
		}
		if len(is) > 2 {
			panic("algorithm requires no border matches more thant two tiles")
		}
		for _, i := range is {
			counts[i]++
		}
	}

	// find the tiles that only have two matching borders, but because the forward/backward combo are considered,
	// look for four instead of 2 matches.
	answer = 1
	matched := 0
	for i, count := range counts {
		if count == 4 {
			matched++
			answer *= tiles[i].ID
		}
	}
	if matched != 4 {
		return 0, fmt.Errorf("did not detect 4 corner tiles: %v", matched)
	}
	return answer, nil
}

var re = regexp.MustCompile(`Tile (\d+):`)

func parseTileNum(s string) (int, error) {
	match := re.FindStringSubmatch(s)
	if match == nil {
		return 0, fmt.Errorf("bad tile header: %v", s)
	}
	return strconv.Atoi(match[1])
}

// https://stackoverflow.com/a/10030772
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
