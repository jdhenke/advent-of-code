package day6

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (ans int, err error) {
	return day6(r, 80)
}

func Part2(r io.Reader) (ans int, err error) {
	return day6(r, 256)
}

func day6(r io.Reader, days int) (ans int, err error) {
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
	total := 0
	for _, x := range nums {
		total += solve(x, days)
	}
	return total, nil
}

type entry struct {
	x, y int
}

var memo = make(map[entry]int)

func solve(timer, days int) (ans int) {
	if ans, ok := memo[entry{timer, days}]; ok {
		return ans
	}
	defer func() {
		memo[entry{timer, days}] = ans
	}()
	if days <= timer {
		return 1
	}
	return solve(6, days-timer-1) + solve(8, days-timer-1)
}
