package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	temp := strings.Split(s, " ")
	list := make([]string, 0)
	for i := 0; i < len(temp); i++ {
		if len(temp[i]) == 0 {
			continue
		}
		list = append(list, temp[i])
	}
	reverse(list)
	return strings.Join(list, " ")
}

func reverse(list []string) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}
