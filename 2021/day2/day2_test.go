package day2_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day2"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	tester.New(t, day2.Part1).Run(
		tester.FromString(testData).Want(150),
		tester.FromFile("input.txt").Want(1762050),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day2.Part2).Run(
		tester.FromString(testData).Want(900),
		tester.FromFile("input.txt").Want(1855892637),
	)
}
