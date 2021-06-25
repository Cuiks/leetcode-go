package main

import "fmt"

func minCut(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPan(s, j, i-1) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}

func isPan(s string, x, y int) bool {
	for x < y {
		if s[x] != s[y] {
			return false
		}
		x++
		y--
	}
	return true
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(minCut("bb"))
}
