package leetcode

import "fmt"

func min(nums ...int) int {
	var res int
	for i, val := range nums {
		if i == 0 {
			res = val
		} else {
			if val < res {
				res = val
			}
		}
	}
	return res
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1, m+1)
	for i, _ := range dp {
		tmp := make([]int, n+1, n+1)
		for j, _ := range tmp {
			tmp[j] = 0

		}
		dp[i] = tmp
	}

	for i, _ := range dp {
		dp[i][0] = i
	}
	for j, _ := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			bias := 1
			if word1[i-1] == word2[j-1] {
				bias = 0
			}
			dp[i][j] = min(dp[i-1][j-1]+bias,
				dp[i][j-1] + 1,
				dp[i-1][j] + 1)
		}
	}


	return dp[m][n]
}

func Test_minDistance() {
	ret := minDistance("horse", "ros")
	fmt.Println("ret = ", ret)

}
