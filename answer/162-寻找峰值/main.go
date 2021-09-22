package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	left := 0
	right := len(nums) - 1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid-1) < get(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}
