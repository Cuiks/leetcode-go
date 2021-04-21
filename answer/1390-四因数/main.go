package main

import "fmt"

func main() {
	nums := []int{21, 4, 7}
	fmt.Println(sumFourDivisors(nums))
}

func sumFourDivisors(nums []int) int {
	hitSum := 0
	for _, num := range nums {
		curSum := 1 + num
		curCount := 2
		start := 2
		for {
			if start*start > num {
				break
			}
			if num%start == 0 {
				curCount += 1
				curSum += start
				if num/start != start {
					curCount += 1
					curSum += num / start
				}
			}
			if curCount > 4 {
				break
			}
			start++
		}
		if curCount == 4 {
			hitSum += curSum
		}
	}
	return hitSum
}
