package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	// dp[i][j]: i-j回文长度
	// dp[i][j] = dp[i+1][j-1]+2  if s[i]==s[j]
	// dp[i][j] = max(dp[i+1][j], dp[i][j-i]) else
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}
