package day21_test

import (
	"github.com/jdhenke/advent-of-code/2022/day21"
	"github.com/jdhenke/advent-of-code/tester"
	"testing"
)

var testData = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestPart1(t *testing.T) {
	tester.New(t, day21.Part1).Run(
		tester.FromString(testData).Want(152),
		tester.FromFile("input.txt").Want(155708040358220),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day21.Part2).Run(
		tester.FromString(testData).Want(301),
		tester.FromFile("input.txt").Want(3342154812537),
	)
}
