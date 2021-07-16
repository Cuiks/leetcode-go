package main

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	start, end := 0, len(numbers)-1
	for start+1 < end {
		for start+1 < end && numbers[start] == numbers[start+1] {
			start++
		}
		for start+1 < end && numbers[end] == numbers[end-1] {
			end--
		}
		mid := (start + end) / 2
		value := numbers[mid]
		if value < numbers[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if numbers[start] <= numbers[end] {
		return numbers[start]
	}
	return numbers[end]
}

func main() {
	fmt.Println(minArray([]int{1,3,3}))
}
