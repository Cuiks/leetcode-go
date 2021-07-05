package main

import "math"

/*
滑动窗口
*/

// 在字符串 S 里面找出：包含 T 所有字母的最小子串
func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	win := map[byte]int{}

	left := 0
	right := 0
	match := 0
	start := 0
	end := 0
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
				if win[c] == need[c] {
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

// 判断 s2 是否包含 s1 的排列
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
		for right-left >= len(s1) {
			if match == len(need) {
				return true
			}
			c = s2[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	return false
}

// 找到 s 中所有是 p 的字母异位词的子串
func findAnagrams(s string, p string) []int {
	need := map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	win := map[byte]int{}

	result := make([]int, 0)
	left, right := 0, 0
	match := 0
	for right < len(s) {
		c := s[right]
		if need[c] != 0 {
			win[c]++
			if need[c] == win[c] {
				match++
			}
		}
		for right-left >= len(need) {
			if right-left == len(p) && match == len(need) {
				result = append(result, left)
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
	return result
}

// 找出不含有重复字符的 最长子串 的长度
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	win := map[byte]int{}
	res := 1
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		win[c]++
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {

}
