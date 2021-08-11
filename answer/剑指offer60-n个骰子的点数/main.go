package main

import "fmt"

var pro = 1.0 / 6.0

func dicesProbability(n int) []float64 {
	dp := make([][]float64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]float64, n*6+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = pro
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= i*6; j++ {
			dp[i][j] = 0
			for k := 1; k <= 6; k++ {
				if j-k < i-1 {
					continue
				}
				dp[i][j] += dp[i-1][j-k] * pro
			}
		}
	}
	res := make([]float64, 0)
	for _, v := range dp[n] {
		if v == 0 {
			continue
		}
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(dicesProbability(10))
}
