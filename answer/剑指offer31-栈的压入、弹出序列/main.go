package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	if len(popped) == 0 {
		return true
	}
	stack := make([]int, 0)
	i, j := 0, 0
	for i < len(popped) && j < len(pushed) {
		stack = append(stack, pushed[j])
		for popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i++
			if i >= len(popped) || len(stack) == 0 {
				break
			}
		}
		j++
	}
	if i == len(popped) {
		return true
	}
	return false
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 0}, []int{1, 0}))
}
