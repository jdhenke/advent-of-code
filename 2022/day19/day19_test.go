package day19_test

import (
	"github.com/jdhenke/advent-of-code/2022/day19"
	"github.com/jdhenke/advent-of-code/tester"
	"testing"
)

var testData = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestPart1(t *testing.T) {
	tester.New(t, day19.Part1).Run(
		tester.FromString(testData).Want(33),
		//tester.FromFile("input.txt").Want(0),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day19.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
