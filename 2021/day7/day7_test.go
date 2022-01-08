package day7_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day7"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `16,1,2,0,4,2,7,1,2,14`

func TestPart1(t *testing.T) {
	tester.New(t, day7.Part1).Run(
		tester.FromString(testData).Want(37),
		tester.FromFile("input.txt").Want(341534),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day7.Part2).Run(
		tester.FromString(testData).Want(168),
		tester.FromFile("input.txt").Want(93397632),
	)
}
