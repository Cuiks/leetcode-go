package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i][j-2]
				if i > 0 && j > 1 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if i == 0 {
					if p[j-1] == '*' {
						dp[i][j] = dp[i][j-1] || dp[i][j-2]
					} else {
						dp[i][j] = false
					}
				} else {
					if p[j-1] == '.' {
						dp[i][j] = dp[i-1][j-1]
					} else if s[i-1] == p[j-1] {
						dp[i][j] = dp[i-1][j-1]
					} else {
						dp[i][j] = false
					}
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
