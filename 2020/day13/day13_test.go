package day13_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day13"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `939
7,13,x,x,59,x,31,19`

func TestPart1(t *testing.T) {
	tester.New(t, day13.Part1).Run(
		tester.FromString(testData).Want(295),
		tester.FromFile("input.txt").Want(2382),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day13.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
