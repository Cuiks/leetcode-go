package main

import (
	"fmt"
	"math"
)

func countDigitOne(n int) int {
	num := map[int]int{}
	num[9] = 1
	i := 9
	for i < n {
		lastI := i
		i = (i+1)*10 - 1
		num[i] = 10*num[lastI] + lastI + 1
	}
	fmt.Println(num)
	res := 0
	if i-n > n-(i/10) {
		index := i / 10
		res = num[index]
		index++
		for index <= n {
			res += count(index)
			index++
		}
	} else {
		index := i
		res = num[index]
		index--
		for index > n {
			res -= count(index)
			index--
		}
	}
	return res
}
func count(i int) int {
	num := 0
	for i > 0 {
		if i%10 == 1 {
			num++
		}
		i = i / 10
	}
	return num
}

func main() {
	fmt.Println(math.Pow(2,32))
	fmt.Println(countDigitOne(824883294))
}
