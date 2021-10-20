package main


// n-1个数+1  等价于  一个数减1
func minMoves(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	ans := 0
	for _, v := range nums {
		ans += v - min
	}
	return ans
}

func main() {

}
