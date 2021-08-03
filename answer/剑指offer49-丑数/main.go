package main

import (
	"fmt"
)

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		v2, v3, v5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(v2, v3), v5)
		if dp[i] == v2 {
			p2++
		}
		if dp[i] == v3 {
			p3++
		}
		if dp[i] == v5 {
			p5++
		}
	}
	return dp[n]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func main() {
	fmt.Println(nthUglyNumber(1690))
}
