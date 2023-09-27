package day20_test

import (
	"github.com/jdhenke/advent-of-code/2022/day20"
	"github.com/jdhenke/advent-of-code/tester"
	"testing"
)

var testData = `1
2
-3
3
-2
0
4`

func TestPart1(t *testing.T) {
	tester.New(t, day20.Part1).Run(
		tester.FromString(testData).Want(3),
		tester.FromFile("input.txt").Want(2622),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day20.Part2).Run(
		tester.FromString(testData).Want(1623178306),
		tester.FromFile("input.txt").Want(1538773034088),
	)
}
