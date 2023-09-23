package input

import (
	"bytes"
	"io"
)

type grid struct {
	data          []byte
	width, height int
	conf          *config
}

var _ Grid = &grid{}

type Grid interface {
	Width() int
	Height() int
	Get(i, j int) string
	Set(i, j int, s string)
	Contains(i, j int) bool
	ForEach(f func(i, j int, s string))
	ForEachAdj(i, j int, f func(s string))
}

type config struct {
	InfiniteRows, InfiniteColumns bool
}

type Option interface {
	apply(c *config)
}

type optionFunc func(c *config)

func (f optionFunc) apply(c *config) {
	f(c)
}

func WithInfiniteRows() Option {
	return optionFunc(func(c *config) {
		c.InfiniteRows = true
	})
}

func WithInfiniteColumns() Option {
	return optionFunc(func(c *config) {
		c.InfiniteColumns = true
	})
}

func NewGrid(r io.Reader, opts ...Option) (Grid, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	conf := &config{
		InfiniteRows:    false,
		InfiniteColumns: false,
	}
	for _, opt := range opts {
		opt.apply(conf)
	}
	data = bytes.TrimSpace(data)
	width := bytes.Index(data, []byte("\n"))
	height := (len(data) + 1) / (width + 1)
	g := &grid{
		data:   data,
		width:  width,
		height: height,
		conf:   conf,
	}
	return g, nil
}

func (g *grid) Width() int {
	if g.conf.InfiniteColumns {
		panic("grid width is infinite")
	}
	return g.width
}

func (g *grid) Height() int {
	if g.conf.InfiniteRows {
		panic("grid height is infinite")
	}
	return g.height
}

func (g *grid) Get(i, j int) string {
	x := g.getIndex(i, j)
	return string(g.data[x : x+1])
}

func (g *grid) Set(i, j int, s string) {
	x := g.getIndex(i, j)
	if len(s) > 1 {
		panic("string is too long: " + s)
	}
	g.data[x] = []byte(s)[0]
}

func (g *grid) ForEach(f func(i, j int, s string)) {
	if g.conf.InfiniteRows || g.conf.InfiniteColumns {
		panic("cannot iterate over an infinite grid")
	}
	for i := 0; i < g.Height(); i++ {
		for j := 0; j < g.Width(); j++ {
			f(i, j, g.Get(i, j))
		}
	}
}

func (g *grid) ForEachAdj(i, j int, f func(s string)) {
	for _, spot := range []struct{ i, j int }{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	} {
		if !g.Contains(spot.i, spot.j) {
			continue
		}
		f(g.Get(spot.i, spot.j))
	}
}

func (g *grid) Contains(i, j int) bool {
	if !g.conf.InfiniteRows && (i < 0 || i >= g.height) {
		return false
	}
	if !g.conf.InfiniteColumns && (j < 0 || j >= g.width) {
		return false
	}
	return true
}

func (g *grid) getIndex(i int, j int) int {
	if g.conf.InfiniteRows {
		i = i % g.height
	}
	if g.conf.InfiniteColumns {
		j = j % g.width
	}
	x := i*(g.width+1) + j
	return x
}
