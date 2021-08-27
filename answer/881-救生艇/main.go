package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	start, end := 0, len(people)-1
	count := 0
	for start <= end {
		count++
		if people[end]+people[start] > limit {
			end--
		} else {
			start++
			end--
		}
	}
	return count
}
func main() {
	fmt.Println(numRescueBoats([]int{3,5,3,4},5))
}
