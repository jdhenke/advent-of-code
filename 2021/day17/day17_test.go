package day17_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day17"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `target area: x=20..30, y=-10..-5`

func TestPart1(t *testing.T) {
	tester.New(t, day17.Part1).Run(
		tester.FromString(testData).Want(45),
		tester.FromFile("input.txt").Want(3003),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day17.Part2).Run(
		tester.FromString(testData).Want(112),
		tester.FromFile("input.txt").Want(940),
	)
}
