package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Strings(dictionary)
	res := ""
	for _, key := range dictionary {
		if valid(s, key) && len(key) > len(res) {
			res = key
		}
	}
	return res
}

func valid(s string, key string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(key) {
		if s[i] == key[j] {
			i++
			j++
		} else {
			i++
		}
	}
	if j == len(key) {
		return true
	}
	return false
}

func main() {
	fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}
