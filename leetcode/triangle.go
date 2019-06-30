package leetcode

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	mini := triangle[len(triangle) - 1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = triangle[i][j] + int(math.Min(float64(mini[j]), float64(mini[j + 1])))
		}
	}
	return mini[0]
}

func Test_Triangle() {
	v := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}
	ret := minimumTotal(v)
	fmt.Println("ret = ", ret)
}
