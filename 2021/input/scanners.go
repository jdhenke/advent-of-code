package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ForEachLine(r io.Reader, f func(line string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		if err := f(s.Text()); err != nil {
			return err
		}
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("error scanning input: %v", err)
	}
	return nil
}

func ForEachInt(r io.Reader, f func(x int)) error {
	return ForEachLine(r, func(line string) error {
		n, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error converting '%s' to int: %v", line, err)
		}
		f(n)
		return nil
	})
}

func ForEachCommand(r io.Reader, f func(cmd string, val int)) error {
	return ForEachLine(r, func(line string) error {
		vals := strings.Split(line, " ")
		if len(vals) != 2 {
			return fmt.Errorf("bad number of values in line: %v", line)
		}
		cmd, valStr := vals[0], vals[1]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return fmt.Errorf("bad value %v: %v", valStr, err)
		}
		f(cmd, val)
		return nil
	})
}
