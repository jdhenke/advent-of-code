package day12_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day12"
	"github.com/jdhenke/advent-of-code/tester"
)

var test1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var test2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var test3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestPart1(t *testing.T) {
	tester.New(t, day12.Part1).Run(
		tester.FromString(test1).Want(10),
		tester.FromString(test2).Want(19),
		tester.FromString(test3).Want(226),
		tester.FromFile("input.txt").Want(4495),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day12.Part2).Run(
		tester.FromString(test1).Want(36),
		tester.FromString(test2).Want(103),
		tester.FromString(test3).Want(3509),
		tester.FromFile("input.txt").Want(131254),
	)
}
