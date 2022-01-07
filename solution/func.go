package solution

import (
	"io"
)

type Func func(r io.Reader) (ans int, err error)
