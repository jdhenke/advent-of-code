package day21_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day21"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `Player 1 starting position: 4
Player 2 starting position: 8`

func TestPart1(t *testing.T) {
	tester.New(t, day21.Part1).Run(
		tester.FromString(testData).Want(739785),
		tester.FromFile("input.txt").Want(903630),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day21.Part2).Run(
		tester.FromString(testData).Want(444356092776315),
		tester.FromFile("input.txt").Want(303121579983974),
	)
}
