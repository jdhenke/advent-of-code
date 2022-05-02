package day22_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day22"
	"github.com/jdhenke/advent-of-code/tester"
)

const testData = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestPart1(t *testing.T) {
	tester.New(t, day22.Part1).Run(
		tester.FromString(testData).Want(306),
		tester.FromFile("input.txt").Want(32272),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day22.Part2).Run(
		tester.FromString(testData).Want(291),
		tester.FromFile("input.txt").Want(33206),
	)
}
