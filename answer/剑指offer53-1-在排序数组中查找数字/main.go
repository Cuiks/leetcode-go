package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end int, target int) int {
	if start >= end {
		if nums[start] == target {
			return 1
		}
		return 0
	}
	mid := (start + end) / 2
	res := 0
	if nums[mid] == target {
		index := mid
		for index <= len(nums)-1 && nums[index] == target {
			res++
			index++
		}
		index = mid - 1
		for index >= 0 && nums[index] == target {
			res++
			index--
		}
		return res
	} else if target < nums[mid] {
		return binarySearch(nums, start, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, end, target)
	}
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1)
	nums = append(nums, 1)
	fmt.Println(search(nums, 1))
}
