package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	index := 0
	start := 0
	indexDict := map[int]int{}
	indexDict[index] = start
	for index < n {
		start = start*10 + 9
		index = 9*(start+1)/10*len(strconv.Itoa(start)) + index
		indexDict[index] = start
	}
	index = index / 10
	start = indexDict[index]
	for index+len(strconv.Itoa(start)) < n {
		index += len(strconv.Itoa(start))
		start++
	}

	startStr := strconv.Itoa(start)
	fmt.Println(startStr, n, index)
	atoi, _ := strconv.Atoi(startStr[n-index : n-index+1])
	return atoi
}

func main() {
	fmt.Println(findNthDigit(1000000000))

}
