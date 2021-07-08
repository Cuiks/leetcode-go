package main

import (
	"fmt"
)

func numSquares(n int) int {
	res := make([]int, n+1)
	res[1] = 1
	if n == 1 {
		return res[1]
	}
	nums := map[int]bool{}
	for i := 1; i <= n/2+1; i++ {
		if i*i > n {
			break
		}
		nums[i*i] = true
	}
	for i := 2; i <= n; i++ {
		res[i] = i
		if _, ok := nums[i]; ok {
			res[i] = 1
			continue
		}
		for j := 1; j*j <= i; j++ {
			if _, ok := nums[j*j]; ok {
				res[i] = min(res[i], res[i-j*j]+1)
			}
		}
	}
	return res[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(7217))
}
