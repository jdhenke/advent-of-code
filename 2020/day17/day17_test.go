package day17_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day17"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `.#.
..#
###`

func TestPart1(t *testing.T) {
	tester.New(t, day17.Part1).Run(
		tester.FromString(testData).Want(112),
		tester.FromFile("input.txt").Want(448),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day17.Part2).Run(
		tester.FromString(testData).Want(848),
		tester.FromFile("input.txt").Want(2400),
	)
}
