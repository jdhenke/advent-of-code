package input

import (
	"io"
	"strconv"
)

func GetNumMatrix(r io.Reader) ([][]int, error) {
	var nums [][]int
	if err := ForEachLine(r, func(line string) error {
		var row []int
		for i := 0; i < len(line); i++ {
			d, err := strconv.Atoi(line[i : i+1])
			if err != nil {
				return err
			}
			row = append(row, d)
		}
		nums = append(nums, row)
		return nil
	}); err != nil {
		return nil, err
	}
	return nums, nil
}
