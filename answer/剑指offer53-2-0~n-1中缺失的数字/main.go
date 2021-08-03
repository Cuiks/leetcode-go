package main

import "fmt"

func missingNumber(nums []int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return nums[i-1] + 1
}

func main() {
	fmt.Println(missingNumber([]int{0, 1}))
}
