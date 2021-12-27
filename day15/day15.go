package day15

import (
	"advent-of-code/input"
	"container/heap"
	"io"
)

type item struct {
	i, j, cost int
}

var _ heap.Interface = &queue{}

type queue []item

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].cost < q[j].cost
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(item))
}

func (q *queue) Pop() interface{} {
	n := len(*q)
	item := (*q)[n-1]
	*q = (*q)[:n-1]
	return item
}

func Part1(r io.Reader) (ans int, err error) {
	nums, err := input.GetNumMatrix(r)
	if err != nil {
		return 0, err
	}
	var q queue
	heap.Push(&q, item{
		i:    1,
		j:    0,
		cost: nums[1][0],
	})
	heap.Push(&q, item{
		i:    0,
		j:    1,
		cost: nums[0][1],
	})
	type entry struct {
		i, j int
	}
	costs := map[entry]int{
		entry{0, 0}: 0,
	}

	neighbors := func(e entry) []entry {
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

	for q.Len() > 0 {
		x := heap.Remove(&q, 0).(item)
		e := entry{x.i, x.j}
		if _, ok := costs[e]; ok {
			continue
		}
		costs[e] = x.cost
		// add neighbors that haven't been hit yet
		for _, n := range neighbors(e) {
			if _, ok := costs[n]; ok {
				continue
			}
			heap.Push(&q, item{
				i:    n.i,
				j:    n.j,
				cost: x.cost + nums[n.i][n.j],
			})
			heap.Init(&q)
		}
	}
	return costs[entry{len(nums) - 1, len(nums[0]) - 1}], nil
}

func Part2(r io.Reader) (ans int, err error) {
	panic("todo")
}
