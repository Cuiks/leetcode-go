package main

func main() {
	nums := []int{0, 1, 1, 1, 2, 2, 3, 3, 4}
	print(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
