package leetcode

import "fmt"

func myPow(x float64, n int) float64 {
	if (n == 0) {
		return 1
	}
	if (x == 0 || x == 1) {
		return x
	}
	if (x == -1) {
		if n%2 == 1 {
			return x
		} else {
			return x * x
		}
	}
	if (n < 0) {
		n = -n
		x = 1 / x
	}

	var pow float64 = 1
	for (n != 0) {
		if (n&1 != 0) {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

func Test_myPow() {
	fmt.Println("test my pow = ", myPow(2, 10))
}
