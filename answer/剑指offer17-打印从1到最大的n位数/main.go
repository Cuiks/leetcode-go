package main

import (
	"fmt"
	"strconv"
)

func printNumbers(n int) []int {
	maxStr := ""
	for i := 0; i < n; i++ {
		maxStr += "9"
	}
	maxInt, _ := strconv.Atoi(maxStr)
	res := make([]int, 0)
	for i := 1; i <= maxInt; i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	fmt.Println(printNumbers(1))
}
