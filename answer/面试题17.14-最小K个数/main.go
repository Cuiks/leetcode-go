package main

import "fmt"

func smallestK(arr []int, k int) []int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		sink(arr, i, len(arr))
	}

	for i := len(arr) - 1; i >= len(arr)-k; i-- {
		swap(arr, 0, i)
		sink(arr, 0, i)
	}
	return arr[len(arr)-k:]
}

func sink(nums []int, i, j int) {
	for {
		left := i*2 + 1
		right := i*2 + 2
		idx := i
		if left < j && nums[left] < nums[idx] {
			idx = left
		}
		if right < j && nums[right] < nums[idx] {
			idx = right
		}
		if idx == i {
			break
		}
		swap(nums, idx, i)
		i = idx
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(smallestK([]int{2, 1, 3, 4, 5, 6, 7, 8, 9, 0}, 5))
}
