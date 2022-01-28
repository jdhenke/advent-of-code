package day13_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day13"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `939
7,13,x,x,59,x,31,19`

func TestPart1(t *testing.T) {
	tester.New(t, day13.Part1).Run(
		tester.FromString(testData).Want(295),
		tester.FromFile("input.txt").Want(2382),
	)
}

var (
	testData2 = `0
17,x,13,19`
	testData3 = `0
67,7,59,61`
	testData4 = `0
67,x,7,59,61`
	testData5 = `0
67,7,x,59,61`
	testData6 = `0
1789,37,47,1889`
)

func TestPart2(t *testing.T) {
	tester.New(t, day13.Part2).Run(
		tester.FromString(testData).Want(1068781),
		tester.FromString(testData2).Want(3417),
		tester.FromString(testData3).Want(754018),
		tester.FromString(testData4).Want(779210),
		tester.FromString(testData5).Want(1261476),
		tester.FromString(testData6).Want(1202161486),
		tester.FromFile("input.txt").Want(906332393333683),
	)
}
