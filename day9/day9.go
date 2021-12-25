package day9

import (
	"advent-of-code/input"
	"io"
	"sort"
	"strconv"
)

func Part1(r io.Reader) (ans int, err error) {
	nums, err := getNums(r)
	if err != nil {
		return 0, err
	}
	total := 0
	points := lowPoints(nums)
	for _, e := range points {
		total += nums[e.i][e.j] + 1
	}
	return total, nil
}

func Part2(r io.Reader) (ans int, err error) {
	nums, err := getNums(r)
	if err != nil {
		return 0, err
	}
	points := lowPoints(nums)
	var sizes []int
	adjacent := func(e entry) []entry {
		var out []entry
		if e.i > 0 {
			out = append(out, entry{e.i - 1, e.j})
		}
		if e.j > 0 {
			out = append(out, entry{e.i, e.j - 1})
		}
		if e.i < len(nums)-1 {
			out = append(out, entry{e.i + 1, e.j})
		}
		if e.j < len(nums[0])-1 {
			out = append(out, entry{e.i, e.j + 1})
		}
		return out
	}
	basinSize := func(i, j int) int {
		covered := map[entry]bool{entry{i, j}: true}
		toCover := []entry{entry{i, j}}
		for len(toCover) > 0 {
			e := toCover[0]
			toCover = toCover[1:]
			for _, a := range adjacent(e) {
				if nums[a.i][a.j] != 9 && !covered[a] {
					covered[a] = true
					toCover = append(toCover, a)
				}
			}
		}
		return len(covered)
	}
	for _, e := range points {
		sizes = append(sizes, basinSize(e.i, e.j))
	}
	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3], nil
}

type entry struct {
	i, j int
}

func getNums(r io.Reader) ([][]int, error) {
	var nums [][]int
	if err := input.ForEachLine(r, func(line string) error {
		var row []int
		for i := 0; i < len(line); i++ {
			d, err := strconv.Atoi(line[i : i+1])
			if err != nil {
				return err
			}
			row = append(row, d)
		}
		nums = append(nums, row)
		return nil
	}); err != nil {
		return nil, err
	}
	return nums, nil
}

func lowPoints(nums [][]int) []entry {
	isLowPoint := func(i, j int) bool {
		if i > 0 && nums[i-1][j] <= nums[i][j] {
			return false
		}
		if j > 0 && nums[i][j-1] <= nums[i][j] {
			return false
		}
		if i < len(nums)-1 && nums[i+1][j] <= nums[i][j] {
			return false
		}
		if j < len(nums[0])-1 && nums[i][j+1] <= nums[i][j] {
			return false
		}
		return true
	}
	var entries []entry
	for i := range nums {
		for j := range nums[i] {
			if isLowPoint(i, j) {
				entries = append(entries, entry{i, j})
			}
		}
	}
	return entries
}
