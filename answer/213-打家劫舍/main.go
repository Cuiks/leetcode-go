package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	left := cal(nums[:length-1])
	right := cal(nums[1:])
	return max(left, right)
}

func cal(nums []int) int {
	length := len(nums)
	numsTotal := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			numsTotal[i] = nums[i]
		} else if i == 1 {
			numsTotal[i] = max(nums[0], nums[1])
		} else {
			numsTotal[i] = max(numsTotal[i-1], numsTotal[i-2]+nums[i])
		}
	}
	return numsTotal[length-1]
}

func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}
