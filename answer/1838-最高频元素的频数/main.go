package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	max := 1
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
}
