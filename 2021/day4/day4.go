package day4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

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

func Part2(r io.Reader) (ans int, err error) {
	numbers, boards, err := parseSetup(r)
	if err != nil {
		return 0, err
	}
	winners := 0
	for _, x := range numbers {
		for j, b := range boards {
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
			if j == 1 {
			}
		}
	}
	return 0, fmt.Errorf("no board won")
}
