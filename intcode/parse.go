package intcode

import (
	"bytes"
	"io"
	"strconv"
)

func Parse(r io.Reader) (Code, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var out Code
	for _, s := range bytes.Split(bytes.TrimSpace(b), []byte{','}) {
		x, err := strconv.Atoi(string(s))
		if err != nil {
			return nil, err
		}
		out = append(out, x)
	}
	return out, nil
}
