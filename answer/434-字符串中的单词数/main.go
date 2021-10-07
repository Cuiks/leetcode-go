package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	split := strings.Split(s, " ")
	splitClean := make([]string, 0)
	for _, v := range split {
		if len(v) != 0 {
			splitClean = append(splitClean, v)
		}
	}
	return len(splitClean)
}

func main() {
	res := countSegments("     ")
	fmt.Println(res)
}
