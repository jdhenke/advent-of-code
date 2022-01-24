package day5_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day5"
	"github.com/jdhenke/advent-of-code/tester"
)

var (
	testData1 = `BFFFBBFRRR`
	testData2 = `FFFBBBFRRR`
	testData3 = `BBFFBBFRLL`
)

func TestPart1(t *testing.T) {
	tester.New(t, day5.Part1).Run(
		tester.FromString(testData1).Want(567),
		tester.FromString(testData2).Want(119),
		tester.FromString(testData3).Want(820),
		tester.FromFile("input.txt").Want(871),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day5.Part2).Run(
		//tester.FromString(testData).Want(0),
		tester.FromFile("input.txt").Want(0),
	)
}
