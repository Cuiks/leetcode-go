package main

import "fmt"

// 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	l := len(triangle)
	f := make([][]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			} else {
				f[i][j] = triangle[i][j]
			}
		}
	}

	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 < 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j >= len(triangle[i-1]) {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
			}
		}
	}
	result := f[l-1][0]
	for i := 0; i < len(f[l-1]); i++ {
		result = min(result, f[l-1][i])
	}
	return result
}

/*
矩阵类型
*/

// 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i] + grid[0][i-1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 不同路径
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

// 不同路径2
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 || f[i-1][0] == 0 {
			f[i][0] = 0
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 || f[0][j-1] == 0 {
			f[0][j] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}

/*
序列类型
*/
// 爬楼梯
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 跳跃游戏
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}

// 跳跃游戏2
func jump(nums []int) int {
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		f[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = min(f[j]+1, f[i])
			}
		}
	}
	return f[len(nums)-1]
}

// 跳跃游戏2 方法2
func jump2(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = 0
	for i := 1; i < n; i++ {
		idx := 0
		for idx <= i && idx+nums[idx] < i {
			idx++
		}
		f[i] = f[idx] + 1
	}
	return f[n-1]
}

// 分割回文串2
func minCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	f := make([]int, len(s)+1)
	f[0] = -1
	f[1] = 0
	for i := 1; i <= len(s); i++ {
		f[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(s)]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	result := f[0]
	for i := 1; i < len(f); i++ {
		result = max(result, f[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 单词拆分
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	maxL, dict := maxLen(wordDict)
	for i := 1; i <= len(s); i++ {
		l := 0
		if i > maxL {
			l = i - maxL
		}
		for j := l; j < i; j++ {
			if f[j] && inDict(s[j:i], dict) {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

func maxLen(wordDict []string) (int, map[string]bool) {
	maxL := 0
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
		if len(v) > maxL {
			maxL = len(v)
		}
	}
	return maxL, dict
}

func inDict(s string, dict map[string]bool) bool {
	_, ok := dict[s]
	return ok
}

// 最长公共子序列
func longestCommonSubsequence(a string, b string) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i <= len(a)+1; i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if a[i] == a[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(a)][len(b)]
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 背包问题
func backPack(m int, A []int) int {
	dp := make([][]bool, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
		}
	}
	for i := m; i >= 0; i-- {
		if dp[len(A)][i] {
			return i
		}
	}
	return 0
}

// 背包问题2
func backPackII(m int, A []int, V []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0
	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j-A[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}
	return dp[len(A)][m]
}

func main() {
	dp := make([][]int, 5)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 6)
	}
	fmt.Println(dp)
}
