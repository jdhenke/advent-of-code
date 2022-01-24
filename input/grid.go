package input

import (
	"bytes"
	"io"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(r)
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
	if g.conf.InfiniteRows {
		i = i % g.height
	}
	if g.conf.InfiniteColumns {
		j = j % g.width
	}
	x := i*(g.width+1) + j
	return string(g.data[x : x+1])
}
