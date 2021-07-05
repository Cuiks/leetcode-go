package main

import "sort"

// 获取无重复数组所有子集
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

// 获取有重复数组所有子集
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	sort.Ints(nums)
	backtrack2(nums, 0, list, &result)
	return result
}

func backtrack2(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, 0)
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}

// 获取全排列
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := map[int]bool{}
	backtrack3(nums, visited, list, &result)
	return result
}

func backtrack3(nums []int, visited map[int]bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack3(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

// 获取有重复元素的全排列
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := map[int]bool{}
	sort.Ints(nums)
	backtrack4(nums, visited, list, &result)
	return result
}

func backtrack4(nums []int, visited map[int]bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(nums))
		copy(ans, list)
		*result = append(*result, ans)
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
		backtrack4(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

func main() {

}
