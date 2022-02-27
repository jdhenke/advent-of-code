package day21_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day21"
	"github.com/jdhenke/advent-of-code/tester"
)

const testData = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestPart1(t *testing.T) {
	tester.New(t, day21.Part1).Run(
		tester.FromString(testData).Want(5),
		tester.FromFile("input.txt").Want(2324),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day21.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
