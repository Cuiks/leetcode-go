package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	indexE := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return false
		}
		if s[i] == 'e' || s[i] == 'E' {
			indexE = i
			break
		}
	}
	if indexE == -1 {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return false
		}
	} else if indexE == 0 {
		return false
	} else if indexE == len(s)-1 {
		return false
	} else {
		left := s[:indexE]
		right := s[indexE+1:]
		_, err := strconv.ParseFloat(left, 64)
		if err != nil {
			return false
		}
		_, err = strconv.Atoi(right)
		if err != nil {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isNumber(". 1"))
}
