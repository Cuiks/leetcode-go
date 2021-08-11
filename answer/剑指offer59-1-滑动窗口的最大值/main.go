package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	queue := make([]int, 0)
	i := 0
	for ; i < k; i++ {
		// 单调递减队列
		queue = push(queue, nums, i)
	}
	res = append(res, nums[queue[0]])
	for ; i < len(nums); i++ {
		queue = push(queue, nums, i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}

func push(queue []int, nums []int, i int) []int {
	for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
		queue = queue[:len(queue)-1]
	}
	queue = append(queue, i)
	return queue
}

func main() {
	fmt.Println(maxSlidingWindow([]int{}, 0))
}
