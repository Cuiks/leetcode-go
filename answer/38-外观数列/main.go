package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := countAndSay(n - 1)
	result := &strings.Builder{}
	cur := prev[0]
	count := 1
	for i := 1; i <= len(prev); i++ {
		var temp byte
		if i < len(prev) {
			temp = prev[i]
		} else {
			temp = ' '
		}

		if cur == temp {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(cur)
			count = 1
			cur = temp
		}
	}
	return result.String()
}

func main() {
	fmt.Println(countAndSay(5))
}
