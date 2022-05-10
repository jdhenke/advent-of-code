package day2_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2019/day2"
	"github.com/jdhenke/advent-of-code/tester"
)

func TestPart1(t *testing.T) {
	tester.New(t, day2.Part1).Run(
		tester.FromFile("input.txt").Want(4090701),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day2.Part2).Run(
		tester.FromFile("input.txt").Want(6421),
	)
}
