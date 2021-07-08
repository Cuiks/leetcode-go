package main

import "fmt"

func isValid(s string) bool {
	list := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ']', '}', ')':
			list = append(list, s[i])
		case '[':
			if len(list) == 0 || list[len(list)-1] != ']' {
				return false
			}
			list = list[:len(list)-1]
		case '{':
			if len(list) == 0 || list[len(list)-1] != '}' {
				return false
			}
			list = list[:len(list)-1]
		case '(':
			if len(list) == 0 || list[len(list)-1] != ')' {
				return false
			}
			list = list[:len(list)-1]
		}
	}
	if len(list) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
