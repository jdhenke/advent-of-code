package day19_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day19"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

func TestPart1(t *testing.T) {
	tester.New(t, day19.Part1).Run(
		tester.FromString(testData).Want(2),
		tester.FromFile("input.txt").Want(151),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day19.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
