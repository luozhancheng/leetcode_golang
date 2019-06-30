package leetcode

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	vec := []int{}
	for i := 0; i < len(nums); i++ {
		index := lowerBound(vec, 0, len(vec)-1, nums[i])
		if index == -1 {
			vec = append(vec, nums[i])
		} else {
			vec[index] = nums[i]
		}
	}
	fmt.Println("vec = ", vec)
	return len(vec)
}

func lowerBound(vec []int, begin int, end int, val int) int {
	if len(vec) == 0 {
		return -1
	}
	if begin > end || begin >= len(vec) {
		return -1
	}

	index := begin + (end-begin) / 2

	if vec[index] >= val {
		if begin == end {
			return index
		}
		return lowerBound(vec, 0, index, val)
	} else {

		return lowerBound(vec, index + 1, end, val)
	}
}

func Test_SubSeq() {
	input := []int{18,55,66,2,3,54}
	ret := lengthOfLIS(input)
	fmt.Println("len = ", ret)
}
