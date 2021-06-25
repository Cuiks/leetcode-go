package main

import "fmt"

func optimalRoute(routes [][]int, source, target int) int {
	dict := make(map[int][]int)
	for i, items := range routes {
		for _, v := range items {
			if inDict(dict, v) {
				dict[v] = append(dict[v], i)
			} else {
				dict[v] = []int{i}
			}
		}
	}
	fmt.Println(dict)
	return count(dict, routes, source, target)
}

func inDict(dict map[int][]int, v int) bool {
	_, ok := dict[v]
	return ok
}

func count(dict map[int][]int, routes [][]int, source, target int) int {
	// 两点不可达
	if !inDict(dict, source) || !inDict(dict, target) {
		return -1
	}
	// 两点一车可达
	sourceItems, _ := dict[source]
	targetItems, _ := dict[target]
	for _, i := range sourceItems {
		for _, j := range targetItems {
			if i == j {
				return 1
			}
		}
	}
	// 递归
	result := len(dict)
	for _, i := range sourceItems {
		for _, j := range routes[i] {
			l := count(dict, routes, j, target)
			if l == -1 {
				return -1
			} else {
				result = min(result, l+1)
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	routes := [][]int{{1, 2, 7}, {3, 6, 7}}
	source := 1
	target := 6
	fmt.Println(optimalRoute(routes, source, target))
}
