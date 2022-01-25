package input

import (
	"io"
)

func ForEachBatch(r io.Reader, f func(batch []string) error) error {
	var batch []string
	if err := ForEachLine(r, func(line string) error {
		if line == "" {
			if len(batch) == 0 {
				return nil
			}
			if err := f(batch); err != nil {
				return err
			}
			batch = nil
			return nil
		}
		batch = append(batch, line)
		return nil
	}); err != nil {
		return err
	}
	if len(batch) > 0 {
		return f(batch)
	}
	return nil
}
