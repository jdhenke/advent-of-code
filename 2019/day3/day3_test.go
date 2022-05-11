package day3_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2019/day3"
	"github.com/jdhenke/advent-of-code/tester"
)

var (
	testData1 = `R8,U5,L5,D3
U7,R6,D4,L4`
	testData2 = `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	testData3 = `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
)

func TestPart1(t *testing.T) {
	tester.New(t, day3.Part1).Run(
		tester.FromString(testData1).Want(6),
		tester.FromString(testData2).Want(159),
		tester.FromString(testData3).Want(135),
		tester.FromFile("input.txt").Want(731),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day3.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
