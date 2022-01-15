package day23_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day23"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

func TestPart1(t *testing.T) {
	tester.New(t, day23.Part1).Run(
		tester.FromString(testData).Want(12521),
		tester.FromFile("input.txt").Want(11536),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day23.Part2).Run(
		tester.FromString(testData).Want(44169),
		tester.FromFile("input.txt").Want(55136),
	)
}
