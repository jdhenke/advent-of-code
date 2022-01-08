package day3_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day3"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestPart1(t *testing.T) {
	tester.New(t, day3.Part1).Run(
		tester.FromString(testData).Want(198),
		tester.FromFile("input.txt").Want(4160394),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day3.Part2).Run(
		tester.FromString(testData).Want(230),
		tester.FromFile("input.txt").Want(4125600),
	)
}
