package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maxLen(slice []int, sum int) int {
	if len(slice) == 0 || sum == 0 {
		return -1
	}
	res := -1
	i, j := 0, -1
	tempSum := 0
	length := 0
	for j < len(slice)-1 {
		j++
		length++
		tempSum += slice[j]
		if tempSum == sum {
			res = max(res, length)
		}
		for tempSum >= sum && i <= j {
			tempSum -= slice[i]
			i++
			length--
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	//fmt.Println(maxLen([]int{1, 2, 3, 4, 3, 2, 1}, 6))
	s := ""
	fmt.Scanln(&s)
	slice := strings.SplitN(s, ",", -1)
	sliceInt := make([]int, 0)
	for _, v := range slice {
		vInt, _ := strconv.Atoi(v)
		sliceInt = append(sliceInt, vInt)
	}
	sum := 0
	fmt.Scanln(&sum)
	fmt.Println(maxLen(sliceInt, sum))
}
