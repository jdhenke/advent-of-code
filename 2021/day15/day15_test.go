package day15_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day15"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestPart1(t *testing.T) {
	tester.New(t, day15.Part1).Run(
		tester.FromString(testData).Want(40),
		tester.FromFile("input.txt").Want(604),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day15.Part2).Run(
		tester.FromString(testData).Want(315),
		tester.FromFile("input.txt").Want(2907),
	)
}
