package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for i, value := range a {
		fmt.Println(i, value)
	}
}
