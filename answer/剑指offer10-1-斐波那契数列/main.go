package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	var a, b, c int
	a, b, c = 0, 1, 0

	for i := 2; i <= n; i++ {
		c = (a%1000000007 + b%1000000007) % 1000000007
		a = b
		b = c
	}
	return c
}

func main() {
	fmt.Println(fib(95))

}
