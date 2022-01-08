package day20_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day20"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func TestPart1(t *testing.T) {
	ans, err := day20.Part1(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 35, ans)
}

func TestPart2(t *testing.T) {
	ans, err := day20.Part2(strings.NewReader(testData))
	require.NoError(t, err)
	assert.Equal(t, 3351, ans)
}
