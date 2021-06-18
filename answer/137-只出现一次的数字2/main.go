package main

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			count += (nums[j] >> i) & 1
		}
		result ^= (count % 3) << i
	}
	return result
}

func main() {

}
