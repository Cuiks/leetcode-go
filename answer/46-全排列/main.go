package main

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	visited := map[int]bool{}
	list := make([]int, 0)
	backtrack(nums, visited, list, &result)
	return result
}

func backtrack(nums []int, visited map[int]bool, list []int, result *[][]int) {
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
		backtrack(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

func main() {

}
