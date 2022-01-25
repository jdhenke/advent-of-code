package day6_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day6"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestPart1(t *testing.T) {
	tester.New(t, day6.Part1).Run(
		tester.FromString(testData).Want(11),
		tester.FromFile("input.txt").Want(6930),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day6.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
