package main

import (
	"fmt"
)

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := (end + start) / 2
	res := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	i, j := start, mid+1
	temp := make([]int, 0)
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			res += j - (mid + 1)
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		temp = append(temp, nums[i])
		res += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		temp = append(temp, nums[j])
	}
	for k := start; k <= end; k++ {
		nums[k] = temp[k-start]
	}
	return res
}

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
