package main

import (
	"fmt"
	"math"
)

var Mod = math.Pow(10, 9) + 7

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	} else if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	// s[i-1]   "s[i-2:i]"
	for i := 2; i <= len(s); i++ {
		if s[i-1] == '*' {
			switch s[i-2] {
			case '*':
				dp[i] = dp[i-1]*9 + dp[i-2]*15
			case '1':
				dp[i] += dp[i-1]*9 + dp[i-2]*9
			case '2':
				dp[i] += dp[i-1]*9 + dp[i-2]*6
			default:
				dp[i] = dp[i-1] * 9
			}
		} else if s[i-1] == '0' {
			switch s[i-2] {
			case '1', '2':
				dp[i] = dp[i-2]
			case '*':
				dp[i] = dp[i-2] * 2
			default:
				dp[i] = 0
			}
		} else if '1' <= s[i-1] && s[i-1] <= '6' {
			switch s[i-2] {
			case '1', '2':
				dp[i] = dp[i-1] + dp[i-2]
			case '*':
				dp[i] = dp[i-1] + dp[i-2]*2
			default:
				dp[i] = dp[i-1]
			}
		} else {
			switch s[i-2] {
			case '1', '*':
				dp[i] = dp[i-1] + dp[i-2]
			default:
				dp[i] = dp[i-1]
			}
		}
		dp[i] = dp[i] % int(Mod)
	}
	return dp[len(s)] % int(Mod)
}

func main() {
	fmt.Println("1*6*7*1*9*6*2*9*2*3*3*6*3*2*2*4*7*2*9*6*0*6*4*4*1*6*9*0*5*9*2*5*7*7*0*6*9*7*1*5*5*9*3*0*4*9*2*6*2*5*7*6*1*9*4*5*8*4*7*4*2*7*1*2*1*9*1*3*0*6*")
	fmt.Println(numDecodings("1*6*7*1*9*6*2*9*2*3*3*6*3*2*2*4*7*2*9*6*0*6*4*4*1*6*9*0*5*9*2*5*7*7*0*6*9*7*1*5*5*9*3*0*4*9*2*6*2*5*7*6*1*9*4*5*8*4*7*4*2*7*1*2*1*9*1*3*0*6*"))
}
