package main

import (
	"fmt"
	"sort"
	"strings"
)

func permutation(s string) []string {
	list := make([]string, 0)
	visited := map[int]bool{}
	result := make([]string, 0)
	s = sortString(s)
	backtrack(s, list, visited, &result)
	return result
}

func backtrack(s string, list []string, visited map[int]bool, result *[]string) {
	if len(list) == len(s) {
		*result = append(*result, strings.Join(list, ""))
	}

	for i := 0; i < len(s); i++ {
		key := string(s[i])
		if visited[i] {
			continue
		}
		if i > 0 && s[i] == s[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		list = append(list, key)
		backtrack(s, list, visited, result)
		list = list[:len(list)-1]
		visited[i] = false
	}

}

func sortString(s string) string {
	stringList := strings.Split(s, "")
	sort.Strings(stringList)
	return strings.Join(stringList, "")
}

func main() {
	fmt.Println(permutation("suvyls"))
}
