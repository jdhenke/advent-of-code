package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ScanInt(r io.Reader, f func(x int)) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()
		n, err := strconv.Atoi(text)
		if err != nil {
			return fmt.Errorf("error converting '%s' to int: %v", text, err)
		}
		f(n)
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("error scanning input: %v", err)
	}
	return nil
}
