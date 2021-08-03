package main

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(s); i++ {
		if s[i-1] == ')' && s[i-2] == '(' {
			dp[i] = dp[i-2] + 2
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses("(()()()((()"))
}
