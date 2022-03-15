package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	index := 0
	for len(strs[0]) > index {
		record := strs[0][index]
		temp := true
		for _, str := range strs[1:] {
			if len(str) <= index {
				temp = false
				break
			}
			if str[index] == record {
				temp = true
			} else {
				temp = false
				break
			}
		}
		if temp {
			index++
		} else {
			break
		}
	}
	return strs[0][:index]
}

func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}
