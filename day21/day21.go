package day21

import (
	"advent-of-code/input"
	"io"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`Player \d+ starting position: (\d+)`)

func Part1(r io.Reader) (answer int, err error) {
	var scores []int64 // int64 to help
	locs := make(map[int]int)
	{
		var player int
		if err := input.ForEachLine(r, func(line string) error {
			scores = append(scores, 0)
			loc, err := strconv.Atoi(re.FindStringSubmatch(line)[1])
			if err != nil {
				return err
			}
			locs[player] = loc - 1
			player++
			return nil
		}); err != nil {
			return 0, err
		}
	}
	die := 100
	count := 0
	roll := func() int {
		count++
		die++
		if die > 100 {
			die = 1
		}
		return die
	}
	for player := 0; ; player = (player + 1) % len(scores) {
		v1, v2, v3 := roll(), roll(), roll()
		locs[player] = (locs[player] + v1 + v2 + v3) % 10
		scores[player] += int64(locs[player] + 1)
		if scores[player] >= 1000 {
			return count * int(scores[(player+1)%len(scores)]), nil
		}
	}
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}
