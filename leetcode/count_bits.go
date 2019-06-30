package leetcode

import "fmt"

func countBits(num int) []int {
	ret := make([]int , num + 1, num + 1)
	for i := 1; i <= num; i++ {
		ret[i] = ret[i & (i - 1)] + 1
	}

	return ret
}




func Test_CountBits() {
	ret := countBits(5)
	for _, v := range ret {
		fmt.Println(v)
	}

}