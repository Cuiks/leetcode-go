package main

import "fmt"

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}

func main() {
	fmt.Println(add(1, 3))
}
