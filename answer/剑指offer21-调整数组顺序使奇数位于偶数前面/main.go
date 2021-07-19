package main

import "fmt"

func exchange(nums []int) []int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i]%2 == 1 {
			i++
			j++
			continue
		}
		if nums[j]%2 == 1 {
			swap(i, j, nums)
			i++
		}
		j++
	}
	return nums
}

func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(exchange([]int{2, 4, 6, 3, 5, 7}))
}
