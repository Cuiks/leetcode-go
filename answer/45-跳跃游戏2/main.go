package main

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		idx := 0
		for idx < i && idx+nums[idx] < i {
			idx++
		}
		dp[i] = dp[idx]+1
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {

}
