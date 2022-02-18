package main

import "fmt"

//func findCenter(edges [][]int) int {
//	dict := map[int]int{}
//	for _, item := range edges {
//		dict[item[0]]++
//		dict[item[1]]++
//	}
//
//	hit := len(dict) - 1
//	for k, v := range dict {
//		if v == hit {
//			return k
//		}
//	}
//	return 0
//}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println(findCenter(edges))
}
