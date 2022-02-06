package day18_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day18"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `1 + 2 * 3 + 4 * 5 + 6`

func TestPart1(t *testing.T) {
	tester.New(t, day18.Part1).Run(
		tester.FromString(testData).Want(71),
		tester.FromFile("input.txt").Want(8298263963837),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day18.Part2).Run(
		tester.FromString(testData).Want(231),
		tester.FromString(`1 + (2 * 3) + (4 * (5 + 6))`).Want(51),
		tester.FromString(`2 * 3 + (4 * 5)`).Want(46),
		tester.FromString(`5 + (8 * 3 + 9 + 3 * 4 * 3)`).Want(1445),
		tester.FromString(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`).Want(669060),
		tester.FromString(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`).Want(23340),
		tester.FromFile("input.txt").Want(145575710203332),
	)
}
