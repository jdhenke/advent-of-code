package day23_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day23"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `389125467`

func TestPart1(t *testing.T) {
	tester.New(t, day23.Part1).Run(
		tester.FromString(testData).Want(67384529),
		tester.FromFile("input.txt").Want(45983627),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day23.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
