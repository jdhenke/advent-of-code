package day14_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day14"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestPart1(t *testing.T) {
	tester.New(t, day14.Part1).Run(
		tester.FromString(testData).Want(165),
		tester.FromFile("input.txt").Want(0),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day14.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
