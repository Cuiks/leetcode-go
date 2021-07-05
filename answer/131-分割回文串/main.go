package main

import (
	"fmt"
	"strings"
)

func partition(s string) [][]string {
	result := make([][]string, 0)
	list := make([]string, 0)
	backtrack(s, 0, list, &result)
	return result
}

func backtrack(s string, pos int, list []string, result *[][]string) {
	if pos >= len(s) {
		if strings.Join(list, "") == s {
			ans := make([]string, len(list))
			copy(ans, list)
			*result = append(*result, ans)
		}
	}
	for i := pos; i < len(s); i++ {
		str := s[pos : i+1]
		if isPan(str) {
			list = append(list, str)
		} else {
			continue
		}
		backtrack(s, i+1, list, result)
		list = list[:len(list)-1]
	}
}

func validList(list []string, s string) bool {
	i := 0
	for _, v := range list {
		if s[i:i+len(v)] != v {
			return false
		}
	}
	return true
}

func isPan(s string) bool {
	if len(s) == 0 {
		return false
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
