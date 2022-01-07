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

type Grid struct {
	Nums [][]int
	Rows int
	Cols int
}

func (g *Grid) Get(i, j int) int {
	if i < 0 || i >= g.Rows || j < 0 || j >= g.Cols {
		panic("bad location")
	}
	val := g.Nums[i%len(g.Nums)][j%len(g.Nums[0])]
	for n := 0; n < i/len(g.Nums)+j/len(g.Nums[0]); n++ {
		val++
		if val > 9 {
			val = 1
		}
	}
	return val
}

func getGraph(r io.Reader) (*Grid, error) {
	nums, err := input.GetNumMatrix(r)
	if err != nil {
		return nil, err
	}
	return &Grid{
		Nums: nums,
		Rows: len(nums),
		Cols: len(nums[0]),
	}, nil
}

func Part1(r io.Reader) (ans int, err error) {
	return day15(r, 1)
}

func Part2(r io.Reader) (ans int, err error) {
	return day15(r, 5)
}

func day15(r io.Reader, m int) (ans int, err error) {
	g, err := getGraph(r)
	if err != nil {
		return 0, err
	}
	g.Rows *= m
	g.Cols *= m

	var q queue
	heap.Push(&q, item{
		i:    1,
		j:    0,
		cost: g.Get(1, 0),
	})
	heap.Push(&q, item{
		i:    0,
		j:    1,
		cost: g.Get(0, 1),
	})
	type entry struct {
		i, j int
	}
	costs := map[entry]int{
		{0, 0}: 0,
	}

	neighbors := func(e entry) []entry {
		var out []entry
		if e.i > 0 {
			out = append(out, entry{e.i - 1, e.j})
		}
		if e.j > 0 {
			out = append(out, entry{e.i, e.j - 1})
		}
		if e.i < g.Rows-1 {
			out = append(out, entry{e.i + 1, e.j})
		}
		if e.j < g.Cols-1 {
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
				cost: x.cost + g.Get(n.i, n.j),
			})
			heap.Init(&q)
		}
	}
	return costs[entry{g.Rows - 1, g.Cols - 1}], nil
}
