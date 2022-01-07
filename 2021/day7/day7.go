package day7

import (
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	nums, err := getNums(r)
	if err != nil {
		return 0, err
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	median := nums[len(nums)/2]
	cost := 0
	for _, x := range nums {
		d := x - median
		if d < 0 {
			d = -d
		}
		cost += d
	}
	return cost, nil
}

func getNums(r io.Reader) ([]int, error) {
	var nums []int
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, s := range strings.Split(string(b), ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, x)
	}
	return nums, nil
}

func Part2(r io.Reader) (ans int, err error) {
	nums, err := getNums(r)
	if err != nil {
		return 0, err
	}
	min, max := nums[0], nums[1]
	for _, x := range nums {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	var bestCost int
	for x := min; x <= max; x++ {
		cost := 0
		for _, n := range nums {
			cost += moveCost(x, n)
		}
		if bestCost == 0 || cost < bestCost {
			bestCost = cost
		}
	}
	return bestCost, nil
}

func moveCost(x, n int) int {
	if x == n {
		return 0
	}
	diff := n - x
	if diff < 0 {
		diff = -diff
	}
	return (diff*diff + diff) / 2
}
