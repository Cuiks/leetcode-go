package main

import "math"

func minSteps(n int) int {
	df := make([]int, n+1)
	df[1] = 0
	for i := 2; i <= n; i++ {
		df[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				df[i] = min(df[i], df[j]+i/j)
				df[i] = min(df[i], df[i/j]+j)
			}
		}
	}
	return df[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
