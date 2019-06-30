package leetcode

import (
	"fmt"
)

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := [2][2]int{}
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		x := i % 2
		y := (i - 1) % 2
		dp[x][0] = Max(Max(dp[y][0] * nums[i], dp[y][1] * nums[i]), nums[i])
		dp[x][1] = Min(Min(dp[y][0] * nums[i], dp[y][1] * nums[i]), nums[i])
		res = Max(dp[x][0], res)
	}

	return res
}

func Test_ProductSub() {
	l := []int{2, 3, -2, 4}
	ret := maxProduct(l)
	fmt.Println("ret = ", ret)
}
