package main

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	} else if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 3
	for i := 3; i <= n; i++ {
		c = (a%1000000007 + b%1000000007) % 1000000007
		a = b
		b = c
	}
	return c % 1000000007
}

func main() {
	fmt.Println(numWays(100))
}
