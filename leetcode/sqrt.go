package leetcode

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	var m int = 0
	l := 1
	r := x
	res := 0

	for l <= r {
		m = r - (r - l) / 2
		if m == x / m {
			return m
		} else if m > x / m {
			r = m - 1
		} else {
			l = m + 1
			res = m
		}
	}

	return res;
}



func Test_mySqrt() {
	fmt.Println("ret = ", mySqrt(8))
}