package main

import "fmt"

var phone = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	list := make([]byte, 0)
	backtrack(digits, 0, list, &result)
	return result
}

func backtrack(digits string, index int, list []byte, result *[]string) {
	if len(list) == len(digits) {
		*result = append(*result, string(list))
		return
	}
	key := phone[digits[index]]
	for i := 0; i < len(key); i++ {
		list = append(list, key[i])
		backtrack(digits, index+1, list, result)
		list = list[0 : len(list)-1]
	}
}

func main() {
	fmt.Println(letterCombinations(""))
}
