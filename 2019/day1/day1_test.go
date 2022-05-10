package day1_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2019/day1"
	"github.com/jdhenke/advent-of-code/tester"
)

func TestPart1(t *testing.T) {
	tester.New(t, day1.Part1).Run(
		tester.FromString(`12`).Want(2),
		tester.FromString(`14`).Want(2),
		tester.FromString(`1969`).Want(654),
		tester.FromString(`100756`).Want(33583),
		tester.FromFile("input.txt").Want(3325156),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day1.Part2).Run(
		tester.FromString(`14`).Want(2),
		tester.FromString(`1969`).Want(966),
		tester.FromString(`100756`).Want(50346),
		tester.FromFile("input.txt").Want(4984866),
	)
}
