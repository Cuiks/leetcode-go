package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	l := len(dp)
	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j >= len(triangle[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	result := dp[l-1][0]
	for i := 1; i < len(triangle[l-1]); i++ {
		result = min(result, dp[l-1][i])
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(triangle)
	fmt.Println(minimumTotal(triangle))
}
