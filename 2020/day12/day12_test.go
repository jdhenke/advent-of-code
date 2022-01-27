package day12_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day12"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `F10
N3
F7
R90
F11`

func TestPart1(t *testing.T) {
	tester.New(t, day12.Part1).Run(
		tester.FromString(testData).Want(25),
		tester.FromFile("input.txt").Want(441),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day12.Part2).Run(
		tester.FromString(testData).Want(286),
		tester.FromFile("input.txt").Want(40014),
	)
}
