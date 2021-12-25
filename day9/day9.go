package day9

import (
	"advent-of-code/input"
	"io"
	"strconv"
)

func Part1(r io.Reader) (ans int, err error) {
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
		return 0, err
	}
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
	total := 0
	for i := range nums {
		for j := range nums[i] {
			if isLowPoint(i, j) {
				total += 1 + nums[i][j]
			}
		}
	}
	return total, nil
}
