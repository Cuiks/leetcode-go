package main

import "fmt"

func fib(n int) int {
	return helper(n)
}

var cache = make(map[int]int)

func helper(n int) int {
	if n < 2 {
		return n
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	res := helper(n-1) + helper(n-2)
	cache[n] = res
	fmt.Println(cache)
	return res
}

func main() {
	fmt.Println(fib(5))
}
