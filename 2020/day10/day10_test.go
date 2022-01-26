package day10_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day10"
	"github.com/jdhenke/advent-of-code/tester"
)

const testData1 = `16
10
15
5
1
11
7
19
6
12
4`

const testData2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestPart1(t *testing.T) {
	tester.New(t, day10.Part1).Run(
		tester.FromString(testData1).Want(7*5),
		tester.FromString(testData2).Want(22*10),
		tester.FromFile("input.txt").Want(2380),
	)
}

// func TestPart2(t *testing.T) {
// 	tester.New(t, day10.Part2).Run(
// 		tester.FromString(testData).Want(0),
// 		tester.FromFile("input.txt").Want(0),
// 	)
// }
