package day1_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day1"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `199
200
208
210
200
207
240
269
260
263`

func TestPart1(t *testing.T) {
	tester.New(t, day1.Part1).Run(
		tester.FromString(testData).Want(7),
		tester.FromFile("input.txt").Want(1791),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day1.Part2).Run(
		tester.FromString(testData).Want(5),
		tester.FromFile("input.txt").Want(1822),
	)
}
