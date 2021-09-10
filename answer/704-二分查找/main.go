package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if target > nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
}
