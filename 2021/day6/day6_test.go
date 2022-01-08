package day6_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day6"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `3,4,3,1,2`

func TestPart1(t *testing.T) {
	tester.New(t, day6.Part1).Run(
		tester.FromString(testData).Want(5934),
		tester.FromFile("input.txt").Want(350605),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day6.Part2).Run(
		tester.FromString(testData).Want(26984457539),
		tester.FromFile("input.txt").Want(1592778185024),
	)
}
