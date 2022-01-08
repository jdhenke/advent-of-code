package day9_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day9"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestPart1(t *testing.T) {
	tester.New(t, day9.Part1).Run(
		tester.FromString(testData).Want(15),
		tester.FromFile("input.txt").Want(532),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day9.Part2).Run(
		tester.FromString(testData).Want(1134),
		tester.FromFile("input.txt").Want(1110780),
	)
}
