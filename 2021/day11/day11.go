package day11

import (
	"advent-of-code/input"
	"io"
)

type entry struct {
	i, j int
}

func Part1(r io.Reader) (ans int, err error) {
	if err := day11(r, func(step, flashes int) bool {
		ans = flashes
		return step < 100
	}); err != nil {
		return 0, err
	}
	return ans, nil
}

func Part2(r io.Reader) (ans int, err error) {
	last := 0
	if err := day11(r, func(step, flashes int) bool {
		ans = step
		ok := flashes-last < 100
		last = flashes
		return ok
	}); err != nil {
		return 0, err
	}
	return ans, nil
}

func day11(r io.Reader, proceed func(step int, flashes int) bool) (err error) {
	nums, err := input.GetNumMatrix(r)
	if err != nil {
		return err
	}

	neighbors := func(e entry) []entry {
		var out []entry
		for i := e.i - 1; i <= e.i+1; i++ {
			for j := e.j - 1; j <= e.j+1; j++ {
				if i < 0 || i >= len(nums) || j < 0 || j >= len(nums[0]) || (i == e.i && j == e.j) {
					continue
				}
				out = append(out, entry{i, j})
			}
		}
		return out
	}

	totalFlashes := 0
	for step := 0; proceed(step, totalFlashes); step++ {
		var queue []entry
		flashes := make(map[entry]bool)
		for i := range nums {
			for j := range nums[i] {
				nums[i][j]++
				if nums[i][j] > 9 {
					e := entry{i, j}
					queue = append(queue, e)
					flashes[e] = true
				}
			}
		}
		for len(queue) > 0 {
			e := queue[0]
			queue = queue[1:]
			// for all neighbors not already flashed, increase by 1, if now 10, say it flashed and add to queue
			for _, n := range neighbors(e) {
				if flashes[n] {
					continue
				}
				nums[n.i][n.j]++
				if nums[n.i][n.j] > 9 {
					queue = append(queue, n)
					flashes[n] = true
				}
			}
		}

		// zero out the flashed ones
		for e := range flashes {
			nums[e.i][e.j] = 0
		}
		totalFlashes += len(flashes)
	}
	return nil
}
