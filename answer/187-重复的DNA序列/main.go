package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	dict := map[string]int{}
	start, end := 0, 9
	for end < len(s) {
		dict[s[start:end+1]]++
		start++
		end++
	}
	result := make([]string, 0)
	for k, v := range dict {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
