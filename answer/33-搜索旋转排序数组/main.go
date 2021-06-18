package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] >= nums[mid] {
			if nums[end] >= target && target > nums[mid] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if target == nums[start] {
		return start
	} else if target == nums[end] {
		return end
	}
	return -1
}
func main() {

}
