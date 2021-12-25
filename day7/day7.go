package day7

import (
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	var nums []int
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	for _, s := range strings.Split(string(b), ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		nums = append(nums, x)
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
