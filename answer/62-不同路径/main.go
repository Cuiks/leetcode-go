package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		for j := 0; j < n; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n)
			}
			dp[i][j] = 1
		}
	}
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
func main() {

}
