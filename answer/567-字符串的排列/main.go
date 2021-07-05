package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	win := map[byte]int{}
	left, right := 0, 0
	match := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] != 0 {
			win[c]++
			if need[c] == win[c] {
				match++
			}
		}
		fmt.Println(win)

		for right-left >= len(s1) {
			if right-left == len(s1) && match == len(need) {
				return true
			}
			c = s2[left]
			left++
			if need[c] != 0 {
				if need[c] == win[c] {
					match--
				}
				win[c]--
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
