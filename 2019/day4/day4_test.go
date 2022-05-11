package day4_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2019/day4"
	"github.com/jdhenke/advent-of-code/tester"
)

func TestPart1(t *testing.T) {
	tester.New(t, day4.Part1).Run(
		tester.FromFile("input.txt").Want(454),
	)
}

//func TestPart2(t *testing.T) {
//	tester.New(t, day4.Part2).Run(
//		tester.FromString(testData).Want(0),
//		tester.FromFile("input.txt").Want(0),
//	)
//}
