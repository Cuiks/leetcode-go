package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	f := make([]int, len(numStr)+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= len(numStr); i++ {
		f[i] = f[i-1]
		if numStr[i-2] == '0' {
			continue
		}
		if numStr[i-2:i] <= "25" {
			f[i] += f[i-2]
		}
	}
	return f[len(numStr)]
}

func main() {
	fmt.Println(translateNum(25))
}
