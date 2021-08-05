package main

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < 64; i++ {
		num := 0
		for j := 0; j < len(nums); j++ {
			num += (nums[j] >> i) & 1
		}
		result ^= (num % 3) << i
	}
	return result
}

func main() {

}
