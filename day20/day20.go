package day20

import (
	"bufio"
	"io"
)

type entry struct {
	i, j int
}

type keyType struct {
	step, i, j int
}

type Image struct {
	key  string
	img0 map[entry]bool
	memo map[keyType]bool
}

func (img *Image) Get(step, i, j int) (ans bool) {
	k := keyType{step, i, j}
	if ans, ok := img.memo[k]; ok {
		return ans
	}
	defer func() {
		img.memo[k] = ans
	}()
	if step == 0 {
		e := entry{i, j}
		return img.img0[e]
	}
	num := 0
	for i2 := i - 1; i2 <= i+1; i2++ {
		for j2 := j - 1; j2 <= j+1; j2++ {
			num <<= 1
			if img.Get(step-1, i2, j2) {
				num |= 1
			}
		}
	}
	return img.key[num:num+1] == "#"
}

func Part1(r io.Reader) (answer int, err error) {
	return day20(r, 2)
}

func day20(r io.Reader, steps int) (answer int, err error) {
	s := bufio.NewScanner(r)
	s.Scan()
	key := s.Text()
	s.Scan()
	img0 := make(map[entry]bool)
	minI, maxI, minJ, maxJ := 0, 0, 0, 0
	{
		i := 0
		for s.Scan() {
			line := s.Text()
			for j := 0; j < len(line); j++ {
				if line[j:j+1] == "#" {
					img0[entry{i, j}] = true
					minI = min(minI, i)
					minJ = min(minJ, j)
					maxI = max(maxI, i)
					maxJ = max(maxJ, j)
				}
			}
			i++
		}
		if err := s.Err(); err != nil {
			return 0, err
		}
	}
	img := &Image{
		key:  key,
		img0: img0,
		memo: make(map[keyType]bool),
	}
	total := 0
	for i := minI - steps; i <= maxI+steps; i++ {
		for j := minJ - steps; j <= maxJ+steps; j++ {
			if img.Get(steps, i, j) {
				total ++
			}
		}
	}
	return total, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Part2(r io.Reader) (answer int, err error) {
	return day20(r, 50)
}
