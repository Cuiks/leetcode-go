package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	dp := make([]bool, l+1)
	dp[0] = true
	max, dict := maxLen(wordDict)
	for i := 1; i <= l; i++ {
		m := 0
		if i > max {
			m = i - max
		}
		for j := m; j < i; j++ {
			if dp[j] && inDict(s[j:i], dict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}

func maxLen(wordDict []string) (int, map[string]bool) {
	max := 0
	dict := make(map[string]bool)
	for _, key := range wordDict {
		dict[key] = true
		if len(key) > max {
			max = len(key)
		}
	}
	return max, dict
}

func inDict(s string, dict map[string]bool) bool {
	_, ok := dict[s]
	return ok
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
