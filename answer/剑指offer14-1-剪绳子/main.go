package main

import "fmt"

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1]
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], dp[i-j]*j), j*(i-j))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func main() {
	fmt.Println(cuttingRope(1000))
}
