package day20_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day20"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func TestPart1(t *testing.T) {
	tester.New(t, day20.Part1).Run(
		tester.FromString(testData).Want(35),
		tester.FromFile("input.txt").Want(5432),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day20.Part2).Run(
		tester.FromString(testData).Want(3351),
		tester.FromFile("input.txt").Want(16016),
	)
}
