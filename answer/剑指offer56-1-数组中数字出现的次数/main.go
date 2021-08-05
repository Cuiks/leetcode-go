package main

func singleNumbers(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}

	result := []int{diff, diff}
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if nums[i]&diff == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

func main() {

}
