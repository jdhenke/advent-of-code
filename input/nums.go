package input

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetNumList(r io.Reader) ([]int, error) {
	var nums []int
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, s := range strings.Split(string(b), ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, x)
	}
	return nums, nil
}
