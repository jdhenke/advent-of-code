package day25

import (
	"io"
	"io/ioutil"
	"strings"
)

func Part1(r io.Reader) (answer int, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(b), "\n")
	var board [][]string
	for _, l := range lines {
		var row []string
		for i := 0; i < len(l); i++ {
			row = append(row, l[i:i+1])
		}
		board = append(board, row)
	}
	type entry struct {
		i, j int
	}
	rows := len(board)
	cols := len(board[0])

	step := 0
	for {
		step++
		moved := false
		// east
		for _, h := range []struct {
			char string
			next func(i, j int) entry
		}{
			{
				char: ">",
				next: func(i, j int) entry {
					return entry{i, (j + 1) % cols}
				},
			},
			{
				char: "v",
				next: func(i, j int) entry {
					return entry{(i + 1) % rows, j}
				},
			},
		} {
			moves := make(map[entry]entry)
			for i := range board {
				for j := range board[i] {
					from := entry{i, j}
					to := h.next(i, j)
					if board[from.i][from.j] == h.char && board[to.i][to.j] == "." {
						moves[from] = to
					}
				}
			}
			moved = moved || len(moves) > 0
			for from, to := range moves {
				board[to.i][to.j] = board[from.i][from.j]
				board[from.i][from.j] = "."
			}
		}
		if !moved {
			return step, nil
		}
	}
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}
