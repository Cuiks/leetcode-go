package main

import "fmt"

func destCity(paths [][]string) string {
	dict := map[string]string{}
	for _, item := range paths {
		dict[item[0]] = item[1]
	}

	curr := paths[0][0]
	for {
		if v, ok := dict[curr]; ok {
			curr = v
		} else {
			break
		}
	}
	return curr
}

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
}
