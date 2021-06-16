package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			temp := make([]byte, 0)
			for stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			// pop '['
			stack = stack[:len(stack)-1]
			// pop num
			idx := 0
			for idx < len(stack) && stack[len(stack)-idx-1] >= '0' && stack[len(stack)-idx-1] <= '9' {
				idx++
			}

			count, _ := strconv.Atoi(string(stack[len(stack)-idx:]))
			stack = stack[:len(stack)-idx]

			for j := 0; j < count; j++ {
				for k := len(temp) - 1; k >= 0; k-- {
					stack = append(stack, temp[k])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	s := "100[lee]"
	fmt.Println(decodeString(s))
}
