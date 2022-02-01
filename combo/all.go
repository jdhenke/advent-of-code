package combo

// All calls f with buf for each possible value of buf where each value is one of vals
func All(buf []int, vals []int, f func(x []int)) {
	all(buf, vals, 0, f)
}

func all(buf []int, vals []int, n int, f func(x []int)) {
	if n == len(buf) {
		f(buf)
		return
	}
	for _, v := range vals {
		buf[n] = v
		all(buf, vals, n+1, f)
	}
}
