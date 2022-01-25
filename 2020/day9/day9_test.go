package day9_test

import (
	"io"
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day9"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestPart1(t *testing.T) {
	t.Run("internal", func(t *testing.T) {
		tester.New(t, func(r io.Reader) (ans int, err error) {
			return day9.Part1WithN(r, 5)
		}).Run(
			tester.FromString(testData).Want(127),
		)
	})
	tester.New(t, day9.Part1).Run(tester.FromFile("input.txt").Want(29221323))
}

func TestPart2(t *testing.T) {
	t.Run("internal", func(t *testing.T) {
		tester.New(t, func(r io.Reader) (ans int, err error) {
			return day9.Part2WithN(r, 5)
		}).Run(
			tester.FromString(testData).Want(62),
		)
	})
	tester.New(t, day9.Part2).Run(tester.FromFile("input.txt").Want(4389369))
}
