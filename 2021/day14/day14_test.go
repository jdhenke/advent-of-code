package day14_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day14"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestPart1(t *testing.T) {
	tester.New(t, day14.Part1).Run(
		tester.FromString(testData).Want(1588),
		tester.FromFile("input.txt").Want(2509),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day14.Part2).Run(
		tester.FromString(testData).Want(2188189693529),
		tester.FromFile("input.txt").Want(2827627697643),
	)
}
