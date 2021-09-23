package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 3 || n==1 {
		return true
	}
	if n%3 != 0 || n == 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

func main() {
	fmt.Println(isPowerOfThree(1))
	fmt.Println(0 % 3)
}
