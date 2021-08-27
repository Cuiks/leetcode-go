package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	list := make([]string, 0)
	backtrack(list, n, 0, 0, &result)
	return result
}

func backtrack(list []string, n, open, close int, result *[]string) {
	if len(list) == 2*n {
		*result = append(*result, strings.Join(list, ""))
		return
	}
	if open < n {
		list = append(list, "(")
		backtrack(list, n, open+1, close, result)
		list = list[:len(list)-1]
	}
	if close < open {
		list = append(list, ")")
		backtrack(list, n, open, close+1, result)
		list = list[:len(list)-1]
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
