package day11_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day11"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestPart1(t *testing.T) {
	tester.New(t, day11.Part1).Run(
		tester.FromString(testData).Want(1656),
		tester.FromFile("input.txt").Want(1601),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day11.Part2).Run(
		tester.FromString(testData).Want(195),
		tester.FromFile("input.txt").Want(368),
	)
}
