package main

import (
	"fmt"
)

func main() {
	a := "111111111111111111111111111111111111111111111"
	fmt.Println(numDecodings(a))
}

func numDecodings(s string) int {
	length := len(s)
	f := make([]int, length+1)
	f[0] = 1
	for i := 1; i <= length; i++ {
		if s[i-1] != '0' {
			f[i] = f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[length]
}
