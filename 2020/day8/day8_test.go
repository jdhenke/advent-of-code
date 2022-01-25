package day8_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day8"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestPart1(t *testing.T) {
	tester.New(t, day8.Part1).Run(
		tester.FromString(testData).Want(5),
		tester.FromFile("input.txt").Want(1709),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day8.Part2).Run(
		tester.FromString(testData).Want(8),
		tester.FromFile("input.txt").Want(1976),
	)
}
