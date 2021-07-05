package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	win := map[byte]int{}

	left, right := 0, 0
	start, end := 0, 0
	match := 0
	min := math.MaxInt64
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c = s[left]
			left++
			if need[c] != 0 {
				if need[c] == win[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}

func main() {
	fmt.Println(minWindow("aa","aa"))
}
