package main

import (
	"fmt"
)

func checkValidString(s string) bool {
	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			left--
		} else {
			left++
		}
		if s[len(s)-i-1] == '(' {
			right--
		} else {
			right++
		}
		if left < 0 || right < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkValidString("("))
}
