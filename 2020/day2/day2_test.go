package day2_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day2"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestPart1(t *testing.T) {
	tester.New(t, day2.Part1).Run(
		tester.FromString(testData).Want(2),
		tester.FromFile("input.txt").Want(414),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day2.Part2).Run(
		tester.FromString(testData).Want(1),
		tester.FromFile("input.txt").Want(413),
	)
}
