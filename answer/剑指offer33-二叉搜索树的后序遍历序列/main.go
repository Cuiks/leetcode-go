package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	return valid(0, len(postorder)-1, postorder)
}

func valid(i, j int, postorder []int) bool {
	if i >= j {
		return true
	}
	m := i
	for postorder[m] < postorder[j] {
		m++
	}
	k := m
	for postorder[m] > postorder[j] {
		m++
	}
	return m == j && valid(i, k-1, postorder) && valid(k, j-1, postorder)
}

func main() {
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))

}
