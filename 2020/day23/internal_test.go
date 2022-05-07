package day23

import (
	"io"
	"testing"

	"github.com/jdhenke/advent-of-code/tester"
)

func TestPart1(t *testing.T) {
	tester.New(t, func(r io.Reader) (ans int, err error) {
		return part1(r, 10)
	}).Run(
		tester.FromString(`389125467`).Want(92658374),
	)

}
