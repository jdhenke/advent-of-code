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

//func TestPart2(t *testing.T) {
//	tester.New(t, day18.Part2).Run(
//		tester.FromString(testData).Want(231),
//		tester.FromFile("input.txt").Want(0),
//	)
//}
