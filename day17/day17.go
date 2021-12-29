package day17

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

// target area: x=20..30, y=-10..-5
var re = regexp.MustCompile(`target area: x=(.+)\.\.(.+), y=(.+)\.\.(.+)`)

type target struct {
	x1, x2, y1, y2 int
}

func (t target) contains(x, y int) bool {
	return x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2
}

func Part1(r io.Reader) (ans int, err error) {
	t, err := getTarget(r)
	if err != nil {
		return 0, err
	}
	height, _ := search(t)
	return height, nil
}

func Part2(r io.Reader) (ans int, err error) {
	t, err := getTarget(r)
	if err != nil {
		return 0, err
	}
	_, hits := search(t)
	return hits, nil
}

func getTarget(r io.Reader) (target, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return target{}, err
	}
	m := re.FindStringSubmatch(string(b))
	if m == nil {
		return target{}, fmt.Errorf("bad input")
	}
	x1, x2, y1, y2 := num(m[1]), num(m[2]), num(m[3]), num(m[4])
	if x2 < 0 {
		x1, x2 = -x2, -x1
	}
	t := target{x1, x2, y1, y2}
	return t, nil
}

func search(t target) (height, hits int) {
	for vy := t.y1; vy <= -t.y1; vy++ {
		for vx := 0; vx <= t.x2; vx++ {
			h, hit := simulate(t, vx, vy)
			if hit {
				hits++
				if h > height {
					height = h
				}
			}
		}
	}
	return height, hits
}

func simulate(t target, vx, vy int) (height int, hit bool) {
	x, y := 0, 0
	for y >= t.y1 {
		if y > height {
			height = y
		}
		if t.contains(x, y) {
			return height, true
		}
		dvx := -1
		if vx == 0 {
			dvx = 0
		}
		x, y, vx, vy = x+vx, y+vy, vx+dvx, vy-1
	}
	return 0, false
}

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("bad input")
	}
	return n
}
