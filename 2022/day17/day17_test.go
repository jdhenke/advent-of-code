package day17_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2022/day17"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestPart1(t *testing.T) {
	tester.New(t, day17.Part1).Run(
		tester.FromString(testData).Want(3068),
		tester.FromFile("input.txt").Want(3168),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day17.Part2).Run(
		tester.FromString(testData).Want(1514285714288),
		tester.FromFile("input.txt").Want(1554117647070),
	)
}
