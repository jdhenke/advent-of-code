package day5_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day5"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	tester.New(t, day5.Part1).Run(
		tester.FromString(testData).Want(5),
		tester.FromFile("input.txt").Want(5698),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day5.Part2).Run(
		tester.FromString(testData).Want(12),
		tester.FromFile("input.txt").Want(15463),
	)
}
