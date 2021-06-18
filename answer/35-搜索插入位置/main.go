package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target >= nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	if target <= nums[start] {
		return start
	} else if target <= nums[end] {
		return end
	} else if target > nums[end] {
		return end + 1
	}
	return - 1
}

func main() {

}
