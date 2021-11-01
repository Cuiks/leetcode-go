package main

import "fmt"

func distributeCandies(candyType []int) int {
	half := len(candyType)/2
	dict := map[int]int{}
	for _, v := range candyType {
		dict[v]++
	}
	dictLen := len(dict)
	if dictLen >= half {
		return half
	}
	return dictLen
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
}
