package day16_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day16"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestPart1(t *testing.T) {
	tester.New(t, day16.Part1).Run(
		tester.FromString(testData).Want(71),
		tester.FromFile("input.txt").Want(22977),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day16.Part2).Run(
		tester.FromFile("input.txt").Want(998358379943),
	)
}
