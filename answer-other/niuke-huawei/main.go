package main

import (
	"bufio"
	"fmt"
	"os"
)

func getLen(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || s[i] == ' ' {
			if s[i] == ' ' {
				break
			} else {
				l++
			}
		}
	}
	return l
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(getLen(input))
	var name string
	fmt.Scanln(&name)

}
