package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(candidates, 0, list, target, &result)
	return result

}

func backtrack(nums []int, pos int, list []int, target int, result *[][]int) {
	count := sum(list)
	if count == target {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	if count > target {
		return
	}
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i, list, target, result)
		list = list[0 : len(list)-1]
	}
}

func sum(list []int) int {
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

func main() {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}
