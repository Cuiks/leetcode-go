package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	fill := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			fill++
		} else {
			if nums[i+1]-nums[i] == 0 {
				return false
			}
			if nums[i+1]-nums[i] == 1 {
				continue
			}
			fill -= nums[i+1] - nums[i] - 1
			if fill < 0 {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}
