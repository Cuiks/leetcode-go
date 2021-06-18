package main

import "fmt"

func largestRectangleArea(heights []int) int {
	max := 0
	if len(heights) == 0 {
		return max
	}

	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = Max(max, h*w)
		}
		stack = append(stack, i)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights := []int{1, 1}
	fmt.Println(largestRectangleArea(heights))
}
