package main

import (
	"fmt"
	"math"
)

func countMaxOrSubsets(nums []int) int {
	result := make([]int, 0)
	list := make([]int, 0)
	max := math.MinInt32
	backTrack(nums, 0, list, &result, &max)
	resInt := 0
	for _, v := range result {
		if v == max {
			resInt += 1
		}
	}
	return resInt
}

func backTrack(nums []int, pos int, list []int, result *[]int, max *int) {
	ans := make([]int, len(list))
	copy(ans, list)
	if len(ans) > 0 {
		v := calOr(ans)
		*result = append(*result, v)
		if v > *max {
			*max = v
		}
	}
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backTrack(nums, i+1, list, result, max)
		list = list[0 : len(list)-1]
	}
}

func calOr(nums []int) int {
	result := 0
	for _, v := range nums {
		result |= v
	}
	return result
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))
}
