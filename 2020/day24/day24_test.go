package day24_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2020/day24"
	"github.com/jdhenke/advent-of-code/tester"
)

var testData = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestPart1(t *testing.T) {
	tester.New(t, day24.Part1).Run(
		tester.FromString(testData).Want(10),
		tester.FromFile("input.txt").Want(388),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day24.Part2).Run(
		tester.FromString(testData).Want(2208),
		tester.FromFile("input.txt").Want(4002),
	)
}
