package day10_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day10"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPart1(t *testing.T) {
	tester.New(t, day10.Part1).Run(
		tester.FromString(testData).Want(26397),
		tester.FromFile("input.txt").Want(370407),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day10.Part2).Run(
		tester.FromString(testData).Want(288957),
		tester.FromFile("input.txt").Want(3249889609),
	)
}
