package day25_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day25"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `5764801
17807724`

func TestPart1(t *testing.T) {
	tester.New(t, day25.Part1).Run(
		tester.FromString(testData).Want(14897079),
		tester.FromFile("input.txt").Want(5025281),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day25.Part2).Run(
		tester.FromString(testData).Want(0),
		tester.FromFile("input.txt").Want(0),
	)
}
