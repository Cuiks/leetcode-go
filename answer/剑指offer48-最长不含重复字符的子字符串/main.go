package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j := 0, 1
	record := map[byte]bool{s[0]: true}
	max := 1
	for j < len(s) && i <= j {
		if record[s[j]] {
			record[s[i]] = false
			i++
		}else {
			record[s[j]] = true
			j++
			if j-i > max {
				max = j - i
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
