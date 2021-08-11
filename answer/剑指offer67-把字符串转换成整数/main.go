package main

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	isNegative := false
	if str[0] == '-' {
		isNegative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	n := 0
	for _, ch := range []byte(str) {
		if ch < '0' || ch > '9' {
			break
		}
		ch -= '0'
		n = n*10 + int(ch)
		if isNegative && -1*n <= math.MinInt32 {
			n = math.MinInt32
			return n
		}
		if n > math.MaxInt32 {
			n = math.MaxInt32
			return n
		}
	}
	if isNegative {
		n = -1 * n
	}
	return n
}

func main() {
	fmt.Println(math.MinInt32, math.MaxInt32)
	fmt.Println(strToInt("-2147483648"))
}
