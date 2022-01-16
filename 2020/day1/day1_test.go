package day1_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day1"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `1721
979
366
299
675
1456`

func TestPart1(t *testing.T) {
	tester.New(t, day1.Part1).Run(
		tester.FromString(testData).Want(514579),
		tester.FromFile("input.txt").Want(876459),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day1.Part2).Run(
		tester.FromString(testData).Want(0),
		tester.FromFile("input.txt").Want(0),
	)
}
