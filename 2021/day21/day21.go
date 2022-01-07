package day21

import (
	"advent-of-code/input"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(b), "\n")
	loc1, err := strconv.Atoi(re.FindStringSubmatch(lines[0])[1])
	if err != nil {
		return 0, err
	}
	loc2, err := strconv.Atoi(re.FindStringSubmatch(lines[1])[1])
	if err != nil {
		return 0, err
	}
	p1Wins, p2Wins := simulate(loc1-1, loc2-1, 0, 0, 0)
	if p1Wins > p2Wins {
		return int(p1Wins), nil
	}
	return int(p2Wins), nil
}

type key struct {
	loc1, loc2     int
	score1, score2 int64
	turn           int
}

type ans struct {
	p1Wins, p2Wins int64
}

var memo = make(map[key]ans)

func simulate(loc1, loc2 int, score1, score2 int64, turn int) (p1Wins, p2Wins int64) {
	e := key{loc1, loc2, score1, score2, turn}
	if ans, ok := memo[e]; ok {
		return ans.p1Wins, ans.p2Wins
	}
	defer func() {
		memo[e] = ans{p1Wins, p2Wins}
	}()
	if score1 >= 21 {
		return 1, 0
	} else if score2 >= 21 {
		return 0, 1
	}
	for _, rolls := range [][]int{
		{1, 1, 1},
		{1, 1, 2},
		{1, 1, 3},
		{1, 2, 1},
		{1, 2, 2},
		{1, 2, 3},
		{1, 3, 1},
		{1, 3, 2},
		{1, 3, 3},
		{2, 1, 1},
		{2, 1, 2},
		{2, 1, 3},
		{2, 2, 1},
		{2, 2, 2},
		{2, 2, 3},
		{2, 3, 1},
		{2, 3, 2},
		{2, 3, 3},
		{3, 1, 1},
		{3, 1, 2},
		{3, 1, 3},
		{3, 2, 1},
		{3, 2, 2},
		{3, 2, 3},
		{3, 3, 1},
		{3, 3, 2},
		{3, 3, 3},
	} {
		s := rolls[0] + rolls[1] + rolls[2]
		subLoc1 := loc1
		subLoc2 := loc2
		subScore1 := score1
		subScore2 := score2
		if turn == 0 {
			subLoc1 = (loc1 + s) % 10
			subScore1 = score1 + int64(subLoc1) + 1
		} else {
			subLoc2 = (loc2 + s) % 10
			subScore2 = score2 + int64(subLoc2) + 1
		}
		subTurn := (turn + 1) % 2
		subP1Wins, subP2Wins := simulate(subLoc1, subLoc2, subScore1, subScore2, subTurn)
		p1Wins += subP1Wins
		p2Wins += subP2Wins
	}
	return p1Wins, p2Wins
}
