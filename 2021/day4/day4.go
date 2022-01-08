package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
Part1 Prompt

--- Day 4: Giant Squid ---
You're already almost 1.5km (almost a mile) below the surface of the ocean,
already so deep that you can't see any sunlight. What you can see, however, is
a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers.
Numbers are chosen at random, and the chosen number is marked on all boards on
which it appears. (Numbers may not appear on all boards.) If all numbers in any
row or any column of a board are marked, that board wins. (Diagonals don't
count.)

The submarine has a bingo subsystem to help passengers (currently, you and the
giant squid) pass the time. It automatically generates a random order in which
to draw numbers and a random set of boards (your puzzle input). For example:

    7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

    22 13 17 11  0
     8  2 23  4 24
    21  9 14 16  7
     6 10  3 18  5
     1 12 20 15 19

     3 15  0  2 22
     9 18 13 17  5
    19  8  7 25 23
    20 11 10 24  4
    14 21 16 12  6

    14 21 17 24  4
    10 16 15  9 19
    18  8 23 26 20
    22 11 13  6  5
     2  0 12  3  7

After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no
winners, but the boards are marked as follows (shown here adjacent to each
other to save space):

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are
still no winners:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

Finally, 24 is drawn:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
     8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
     6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
     1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

At this point, the third board wins because it has at least one complete row or
column of marked numbers (in this case, the entire top row is marked: 14 21 17
24 4).

The score of the winning board can now be calculated. Start by finding the sum
of all unmarked numbers on that board; in this case, the sum is 188. Then,
multiply that sum by the number that was just called when the board won, 24, to
get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win
first. What will your final score be if you choose that board?
*/
func Part1(r io.Reader) (ans int, err error) {
	numbers, boards, err := parseSetup(r)
	if err != nil {
		return 0, err
	}
	for _, x := range numbers {
		for _, b := range boards {
			if b.Mark(x) {
				return x * b.Sum, nil
			}
		}
	}
	return 0, fmt.Errorf("no board won")
}

/*
Part2 Prompt

--- Part Two ---
On the other hand, it might be wise to try a different strategy: let the giant
squid win.

You aren't sure how many bingo boards a giant squid could play at once, so
rather than waste time counting its arms, the safe thing to do is to figure out
which board will win last and choose that one. That way, no matter which boards
it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after
13 is eventually called and its middle column is completely marked. If you were
to keep playing until this point, the second board would have a sum of unmarked
numbers equal to 148 for a final score of 148 * 13 = 1924.

Figure out which board will win last. Once it wins, what would its final score
be?
*/
func Part2(r io.Reader) (ans int, err error) {
	numbers, boards, err := parseSetup(r)
	if err != nil {
		return 0, err
	}
	winners := 0
	for _, x := range numbers {
		for _, b := range boards {
			if b.Won {
				continue
			}
			if b.Mark(x) {
				winners++
				if winners == len(boards) {
					return x * b.Sum, nil
				}
				b.Won = true
			}
		}
	}
	return 0, fmt.Errorf("no board won")
}

type Entry struct {
	I, J int
}

type Board struct {
	Nums map[int][]Entry
	Cols [5]int
	Rows [5]int
	Sum  int
	Won  bool
}

func (b *Board) Mark(x int) bool {
	if es, ok := b.Nums[x]; ok {
		for _, e := range es {
			b.Sum -= x
			b.Rows[e.I]++
			if b.Rows[e.I] == 5 {
				return true
			}
			b.Cols[e.J]++
			if b.Cols[e.J] == 5 {
				return true
			}
		}
	}
	return false
}

func parseSetup(r io.Reader) ([]int, []*Board, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, nil, fmt.Errorf("could not scan numbers")
	}
	var nums []int
	for _, s := range strings.Split(s.Text(), ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil, err
		}
		nums = append(nums, x)
	}
	if !s.Scan() {
		return nil, nil, fmt.Errorf("no first board")
	}
	var boards []*Board
	for {
		nums := make(map[int][]Entry)
		sum := 0
		for i := 0; i < 5; i++ {
			s.Scan()
			for j := 0; j < 5; j++ {
				s := s.Text()[j*3 : j*3+2]
				x, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					return nil, nil, err
				}
				nums[x] = append(nums[x], Entry{
					I: i,
					J: j,
				})
				sum += x
			}
		}
		boards = append(boards, &Board{
			Nums: nums,
			Cols: [5]int{},
			Rows: [5]int{},
			Sum:  sum,
		})
		if !s.Scan() {
			break
		}
	}
	return nums, boards, nil
}
