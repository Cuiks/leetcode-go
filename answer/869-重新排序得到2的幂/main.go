package main

import "fmt"

var dict = map[[10]int]bool{}

func count(n int) [10]int {
	res := [10]int{}
	for n > 0 {
		res[n%10]++
		n /= 10
	}
	return res
}

func init() {
	for i := 1; i <= 1e9; i <<= 1 {
		dict[count(i)] = true
	}
}

func reorderedPowerOf2(n int) bool {
	return dict[count(n)]
}

func main() {
	fmt.Println(reorderedPowerOf2(1420))
}
