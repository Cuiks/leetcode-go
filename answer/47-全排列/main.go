package main

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := map[int]bool{}
	sort.Ints(nums)
	backtrack(nums, visited, list, &result)
	return result
}

func backtrack(nums []int, visited map[int]bool, list []int, result *[][]int) {
	if len(nums) == len(list) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

func main() {

}
