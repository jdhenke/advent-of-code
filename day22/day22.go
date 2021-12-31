package day22

import (
	"advent-of-code/input"
	"io"
	"regexp"
	"sort"
	"strconv"
)

// on x=4..48,y=-44..10,z=-45..4
var re = regexp.MustCompile(`(.+) x=(.+)\.\.(.+),y=(.+)\.\.(.+),z=(.+)\.\.(.+)`)

const dims = 3

type Cube struct {
	bounds [dims][2]int
	on     bool
}

func (c *Cube) Intersects(other *Cube) bool {
	for i := 0; i < dims; i++ {
		if !overlap(c.bounds[i], other.bounds[i]) {
			return false
		}
	}
	return true
}

func overlap(r1, r2 [2]int) bool {
	return in(r1[0], r2) || in(r1[1]-1, r2) || in(r2[0], r1) || in(r2[1]-1, r1)
}

func in(x int, r [2]int) bool {
	return x >= r[0] && x < r[1]
}

// Split produces the cubes that represent c with other taken out of it
func (c *Cube) Split(other *Cube) []*Cube {
	var ranges [dims][][2]int // for each dimension there's three ranges
	for i := 0; i < dims; i++ {
		ranges[i] = segment(c.bounds[i], other.bounds[i])
	}
	// combine all d0 options with al
	var f func(i int) [][][2]int
	f = func(i int) [][][2]int {
		if i == dims {
			return [][][2]int{nil}
		}
		var out [][][2]int
		for _, dOpt := range ranges[i] {
			// combine this di option with all other options for i+1 onward
			for _, sub := range f(i + 1) {
				opt := append([][2]int{dOpt}, sub...)
				out = append(out, opt)
			}
		}
		return out
	}
	var out []*Cube
	for _, b := range f(0) {
		var a [3][2]int
		for i := 0; i < dims; i++ {
			a[i] = b[i]
		}
		sub := &Cube{
			bounds: a,
			on:     c.on,
		}
		if other.Intersects(sub) {
			continue
		}
		if !c.Intersects(sub) {
			continue
		}
		out = append(out, sub)
	}
	return out
}

func segment(r0, r1 [2]int) [][2]int {
	xs := []int{r0[0], r0[1], r1[0], r1[1]}
	sort.Ints(xs)
	var out [][2]int
	for _, r := range [][2]int{
		{xs[0], xs[1]},
		{xs[1], xs[2]},
		{xs[2], xs[3]},
	} {
		if r[1] <= r[0] {
			continue
		}
		out = append(out, r)
	}
	return out
}

func Part1(r io.Reader) (answer int, err error) {
	allCubes, err := readCubes(r)
	if err != nil {
		return 0, err
	}
	var cubes []*Cube
	boundary := &Cube{
		bounds: [3][2]int{
			{-50, 51},
			{-50, 51},
			{-50, 51},
		},
	}
	for _, other := range allCubes {
		if !boundary.Intersects(other) {
			continue
		}
		var newCubes []*Cube
		for _, c := range cubes {
			if c.Intersects(other) {
				newCubes = append(newCubes, c.Split(other)...)
			} else {
				newCubes = append(newCubes, c)
			}
		}
		newCubes = append(newCubes, other)
		cubes = newCubes
	}

	count := 0
	for _, c := range cubes {
		if !c.on {
			continue
		}
		prod := 1
		for i := 0; i < dims; i++ {
			prod *= c.bounds[i][1] - c.bounds[i][0]
		}
		count += prod
	}
	return count, nil
}

func readCubes(r io.Reader) ([]*Cube, error) {
	var cubes []*Cube
	if err := input.ForEachLine(r, func(line string) error {
		cubes = append(cubes, parseCube(line))
		return nil
	}); err != nil {
		return nil, err
	}
	return cubes, nil
}

func Part2(r io.Reader) (answer int, err error) {
	return 0, nil
}

func parseCube(line string) *Cube {
	parts := re.FindStringSubmatch(line)
	cube := &Cube{
		bounds: [dims][2]int{
			{num(parts[2]), num(parts[3]) + 1},
			{num(parts[4]), num(parts[5]) + 1},
			{num(parts[6]), num(parts[7]) + 1},
		},
		on: parts[1] == "on",
	}
	return cube
}

func num(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
