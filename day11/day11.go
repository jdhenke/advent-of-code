package day11

import (
	"advent-of-code/input"
	"io"
)

type entry struct {
	i, j int
}

func Part1(r io.Reader) (ans int, err error) {
	nums, err := input.GetNumMatrix(r)
	if err != nil {
		return 0, err
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
	for step := 0; step < 100; step++ {
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
	return totalFlashes, nil
}
