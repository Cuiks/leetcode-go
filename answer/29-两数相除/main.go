package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		} else {
			return -dividend
		}
	}
	flag := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flag = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	result := div(dividend, divisor)
	if flag {
		return -result
	}
	return result
}

func div(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	count := 1
	res := divisor
	for res+res <= dividend {
		count += count
		res += res
	}
	return count + div(dividend-res, divisor)
}

func main() {
	fmt.Println(divide(2147483647, 3))
}
