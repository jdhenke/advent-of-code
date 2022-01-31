package day15_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day15"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `0,3,6`

func TestPart1(t *testing.T) {
	tester.New(t, day15.Part1).Run(
		tester.FromString(testData).Want(436),
		tester.FromFile("input.txt").Want(706),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day15.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
