package day13_test

import (
	"strings"
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day13"
	"github.com/jdhenke/advent-of-code/tester"
	"github.com/stretchr/testify/require"
)

var testData = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestPart1(t *testing.T) {
	tester.New(t, day13.Part1).Run(
		tester.FromString(testData).Want(17),
		tester.FromFile("input.txt").Want(682),
	)
}

func TestPart2(t *testing.T) {
	_, err := day13.Part2(strings.NewReader(testData))
	require.NoError(t, err)
}
