package day20

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlip(t *testing.T) {
	tile := Tile{
		Data: []string{
			"abc",
			"def",
			"ghi",
		},
	}
	flip := tile.flip()
	want := []string{
		"adg",
		"beh",
		"cfi",
	}
	assert.Equal(t, want, flip.Data)
}

func TestRotate(t *testing.T) {
	tile := Tile{
		Data: []string{
			"abc",
			"def",
			"ghi",
		},
	}
	rot := tile.rotate()
	want := []string{
		"gda",
		"heb",
		"ifc",
	}
	assert.Equal(t, want, rot.Data)
}

func TestTile_ForAllBorders(t *testing.T) {
	var got []BorderAndSide
	Tile{
		Data: []string{
			"abc",
			"def",
			"ghi",
		},
	}.ForAllBorders(func(bs BorderAndSide) {
		got = append(got, bs)
	})
	want := []BorderAndSide{
		{
			Side:   SideTop,
			Border: "abc",
		},
		{
			Side:   SideRight,
			Border: "cfi",
		},
		{
			Side:   SideBottom,
			Border: "ghi",
		},
		{
			Side:   SideLeft,
			Border: "adg",
		},
	}
	assert.Equal(t, want, got)
}

const testData = `Tile 2311:
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
..#.###...`

func TestSolve(t *testing.T) {
	grid, err := solve(strings.NewReader(testData))
	require.NoError(t, err)
	want := [][]int{
		{1171, 2473, 3079},
		{1489, 1427, 2311},
		{2971, 2729, 1951},
	}
	for i := 0; i < len(want); i++ {
		for j := 0; j < len(want[i]); j++ {
			assert.Equal(t, want[i][j], grid[i][j].ID, "%d %d", i, j)
		}
	}
}
