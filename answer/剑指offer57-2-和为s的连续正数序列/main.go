package main

import (
	"fmt"
)

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	for i := 2; i < target; i++ {
		start, end := -1, -1
		// 奇数情况
		if target%i == 0 && target/i >= 3 {
			if target/i%2 == 0 {
				continue
			} else {
				k := (target/i - 1) / 2
				start, end = i-k, i+k
			}
		}
		// 偶数情况
		if target%(i+i+1) == 0 {
			k := target/(i+i+1) - 1
			start, end = i-k, i+1+k
		}
		if start > 0 {
			temp := make([]int, 0)
			for start <= end {
				temp = append(temp, start)
				start++
			}
			result = append(result, temp)
		}
	}
	return result
}

func main() {
	fmt.Println(findContinuousSequence(15))
}
