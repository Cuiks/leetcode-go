package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	l := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < l; i++ {
		if obstacleGrid[i][0] != 1 && dp[i-1][0] != 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] != 1 && dp[0][i-1] != 0 {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < l; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[l-1][n-1]
}
func main() {
	grid := [][]int{{0, 1}, {0, 0}}
	fmt.Println(uniquePathsWithObstacles(grid))
}
