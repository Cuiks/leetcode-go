package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	list := make([]int, 0)
	for num != 0 {
		list = append(list, num%7)
		num /= 7
	}
	res := 0
	for i := len(list) - 1; i >= 0; i-- {
		res = res*10 + list[i]
	}
	resStr := strconv.Itoa(res)
	return resStr
}

func main() {
	fmt.Println(convertToBase7(-7))
}
