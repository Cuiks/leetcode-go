package main

import "fmt"

func twoSum(nums []int, target int) []int {
	half := target / 2
	start, end := 0, len(nums)-1
	var mid int
	for start+1 < end {
		mid = (start + end) / 2
		if half < nums[mid] {
			end = mid
		} else {
			start = mid
		}
	}
	for start >= 0 && end <= len(nums)-1 {
		left, right := nums[start], nums[end]
		if left+right == target {
			return []int{left, right}
		} else if left+right < target {
			end++
		} else if left+right > target {
			start--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{16,16,18,24,30,32}, 48))

}
