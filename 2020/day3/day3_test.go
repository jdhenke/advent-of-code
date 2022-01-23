package day3_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day3"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestPart1(t *testing.T) {
	tester.New(t, day3.Part1).Run(
		tester.FromString(testData).Want(7),
		tester.FromFile("input.txt").Want(153),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day3.Part2).Run(
		tester.FromString(testData).Want(0),
		tester.FromFile("input.txt").Want(0),
	)
}
