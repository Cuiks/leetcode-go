package main

import "fmt"

func constructArr(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	left := make([]int, len(a)) // 下三角
	left[0] = 1
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1] * a[i-1]
	}
	res[len(a)-1] = left[len(a)-1]
	temp := 1
	for j := len(a) - 2; j >= 0; j-- {
		temp *= a[j+1]
		res[j] = left[j] * temp
	}
	return res
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
