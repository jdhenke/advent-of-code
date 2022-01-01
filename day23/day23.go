package day23

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sort"
)

// #############
// #01.2.3.4.56#
// ###C#A#B#D### 0
//   #C#A#D#B#   1
//   #########
//    1 2 3 4

type Board struct {
	A1, A2, B1, B2, C1, C2, D1, D2 int
}

func Part1(r io.Reader) (answer int, err error) {
	b, err := parseBoard(r)
	if err != nil {
		return 0, err
	}
	ans, ok := solve(b)
	if !ok {
		return 0, fmt.Errorf("could not solve board")
	}
	return ans, nil
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}

var re = regexp.MustCompile(`#############
#...........#
###(.)#(.)#(.)#(.)###
  #(.)#(.)#(.)#(.)#
  #########`)

func parseBoard(r io.Reader) (Board, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return Board{}, err
	}
	parts := re.FindStringSubmatch(string(b))
	reverse := map[int]string{
		10: parts[1],
		20: parts[2],
		30: parts[3],
		40: parts[4],
		11: parts[5],
		21: parts[6],
		31: parts[7],
		41: parts[8],
	}
	lookup := make(map[string][]int)
	for loc, pieceStr := range reverse {
		lookup[pieceStr] = append(lookup[pieceStr], loc)
		sort.Ints(lookup[pieceStr])
	}
	return Board{
		A1: lookup["A"][0],
		A2: lookup["A"][1],
		B1: lookup["B"][0],
		B2: lookup["B"][1],
		C1: lookup["C"][0],
		C2: lookup["C"][1],
		D1: lookup["D"][0],
		D2: lookup["D"][1],
	}, nil
}

var memo = make(map[Board]IntBool)

type IntBool struct {
	Int  int
	Bool bool
}

func solve(b Board) (ans int, ok bool) {
	if ans, ok := memo[b]; ok {
		return ans.Int, ans.Bool
	}
	defer func() {
		memo[b] = IntBool{ans, ok}
	}()
	if solved(b) {
		return 0, true
	}
	var cheapest int
	for next, moveCost := range getAllNext(b) {
		ans, ok := solve(next)
		if !ok {
			continue
		}
		cost := moveCost + ans
		if cheapest == 0 || cost < cheapest {
			cheapest = cost
		}
	}
	if cheapest == 0 {
		return 0, false
	}
	return cheapest, true
}

func solved(b Board) bool {
	return inHome(1, b.A1, b.A2) && inHome(2, b.B1, b.B2) && inHome(3, b.C1, b.C2) && inHome(4, b.D1, b.D2)
}

func inHome(h int, x1 int, x2 int) bool {
	return x1/10 == h && x2/10 == h
}

func getAllNext(b Board) map[Board]int {
	out := make(map[Board]int)
	for _, x := range []struct {
		loc         *int
		costPerStep int
	}{
		{&b.A1, 1},
		{&b.A2, 1},
		{&b.B1, 10},
		{&b.B2, 10},
		{&b.C1, 100},
		{&b.C2, 100},
		{&b.D1, 1000},
		{&b.D2, 1000},
	} {
		for loc, steps := range moves(b, *x.loc) {
			old := *x.loc
			*x.loc = loc
			out[b] = steps * x.costPerStep
			*x.loc = old
		}
	}
	return out
}

func moves(b Board, pos int) map[int]int {
	out := make(map[int]int)
	locs := map[int]int{
		b.A1: 1,
		b.A2: 1,
		b.B1: 2,
		b.B2: 2,
		b.C1: 3,
		b.C2: 3,
		b.D1: 4,
		b.D2: 4,
	}
	piece := locs[pos]
	try := func(start, end int) {
		steps, n := getPath(start, end)
		for _, loc := range steps {
			if locs[loc] != 0 {
				return
			}
		}
		out[end] = n
	}
	switch pos {
	case 0, 1, 2, 3, 4, 5, 6:
		try(pos, (piece*10)+1)
		if locs[(piece*10)+1] == piece {
			try(pos, piece*10)
		}
	case 10, 11, 20, 21, 30, 31, 40, 41:
		if piece*10+1 == pos {
			return nil
		}
		if piece*10 == pos && locs[piece*10+1] == piece {
			return nil
		}
		for dest := 0; dest <= 6; dest++ {
			try(pos, dest)
		}
	}
	return out
}

func getPath(a, b int) (path []int, n int) {
	if a == b {
		return nil, 0
	}
	if a > b {
		sp, n := getPath(b, a)
		for i := len(sp) - 2; i >= 0; i-- { // ignore a at end
			path = append(path, sp[i])
		}
		path = append(path, b) // add b that the reverse wouldn't
		return path, n
	}
	add := func(next int, cost int) (path []int, n int) {
		sp, sn := getPath(next, b)
		path = append(path, next)
		path = append(path, sp...)
		return path, cost + sn
	}
	if a == 0 {
		return add(1, 1)
	}
	switch a {
	case 0:
		return add(1, 1)
	case 1:
		switch b {
		case 10, 11:
			return add(10, 2)
		}
		return add(2, 2)
	case 2:
		switch b {
		case 10, 11:
			return add(10, 2)
		case 20, 21:
			return add(20, 2)
		}
		return add(3, 2)
	case 3:
		switch b {
		case 10, 11:
			return add(2, 2)
		case 20, 21:
			return add(20, 2)
		case 30, 31:
			return add(30, 2)
		}
		return add(4, 2)
	case 4:
		switch b {
		case 40, 41:
			return add(40, 2)
		case 30, 31:
			return add(30, 2)
		}
		return add(3, 2)
	case 5:
		switch b {
		case 40, 41:
			return add(40, 2)
		}
		return add(4, 2)
	case 6:
		return add(5, 1)
	case 10:
		return add(11, 1)
	case 20:
		return add(21, 1)
	case 30:
		return add(31, 1)
	case 40:
		return add(41, 1)
	}
	panic(fmt.Sprintf("%#v", a))
}
