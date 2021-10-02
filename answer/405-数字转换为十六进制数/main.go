package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	base := "0123456789abcdef"
	ans := ""
	if num < 0 {
		num += 0x100000000
	}

	for num != 0 {
		ans += string(base[num%16])
		num /= 16
	}
	res := revers(ans)
	return res
}

func revers(s string) string {
	a := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		a[i], a[len(s)-i-1] = a[len(s)-i-1], a[i]
	}
	return string(a)
}

func main() {
	fmt.Println(toHex(100000000))
}
