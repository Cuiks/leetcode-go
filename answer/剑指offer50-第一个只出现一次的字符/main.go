package main

import "fmt"

func firstUniqChar(s string) byte {
	dict := map[byte]int{}
	for i := 0; i < len(s); i++ {
		dict[s[i]] += 1
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	dict := map[byte]int{}
	dict['a'] += 1
	fmt.Println(dict['a'])
}
