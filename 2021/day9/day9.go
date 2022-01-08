package day9

import (
	"advent-of-code/input"
	"io"
	"sort"
)

/*
Part1 Prompt

--- Day 9: Smoke Basin ---
These caves seem to be lava tubes. Parts are even still volcanically active;
small hydrothermal vents release smoke into the caves that slowly settles like
rain.

If you can model how the smoke flows through the caves, you might be able to
avoid it and be that much safer. The submarine generates a heightmap of the
floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the
following heightmap:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678

Each number corresponds to the height of a particular location, where 9 is the
highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than
any of its adjacent locations. Most locations have four adjacent locations (up,
down, left, and right); locations on the edge or corner of the map have three
or two adjacent locations, respectively. (Diagonal locations do not count as
adjacent.)

In the above example, there are four low points, all highlighted: two are in
the first row (a 1 and a 0), one is in the third row (a 5), and one is in the
bottom row (also a 5). All other locations on the heightmap have some lower
adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the
risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of
all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk
levels of all low points on your heightmap?
*/
func Part1(r io.Reader) (ans int, err error) {
	nums, err := input.GetNumMatrix(r)
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

/*
Part2 Prompt

--- Part Two ---
Next, you need to find the largest basins so you know what areas are most
important to avoid.

A basin is all locations that eventually flow downward to a single low point.
Therefore, every low point has a basin, although some basins are very small.
Locations of height 9 do not count as being in any basin, and all other
locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the
low point. The example above has four basins.

The top-left basin, size 3:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678

The top-right basin, size 9:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678

The middle basin, size 14:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678

The bottom-right basin, size 9:

    2199943210
    3987894921
    9856789892
    8767896789
    9899965678

Find the three largest basins and multiply their sizes together. In the above
example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
*/
func Part2(r io.Reader) (ans int, err error) {
	nums, err := input.GetNumMatrix(r)
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
		covered := map[entry]bool{{i, j}: true}
		toCover := []entry{{i, j}}
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
