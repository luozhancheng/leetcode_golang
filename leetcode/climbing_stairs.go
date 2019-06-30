package leetcode

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var result int
	var oneStep = 2
	var twoStep = 1

	for i := 2 ; i < n; i++ {
		result = oneStep + twoStep
		twoStep = oneStep
		oneStep = result
	}
	return result;
}



func Run_TestClim() {
	ret := climbStairs(3)
	fmt.Println("ret = ", ret)
}