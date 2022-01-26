package day11_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day11"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestPart1(t *testing.T) {
	tester.New(t, day11.Part1).Run(
		tester.FromString(testData).Want(37),
		tester.FromFile("input.txt").Want(2441),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day11.Part2).Run(
		tester.FromString(testData).Want(26),
		tester.FromFile("input.txt").Want(2190),
	)
}
