package main

import "fmt"

func findMinMoves(machines []int) int {
	total := sum(machines)
	if total%len(machines) != 0 {
		return -1
	}

	avg := total / len(machines)
	sum := 0
	ans := 0
	for _, num := range machines {
		num -= avg
		sum += num
		ans = max(ans, max(abs(sum), num))
	}
	return ans
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func sum(list []int) int {
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

func main() {
	fmt.Println(findMinMoves([]int{4, 1, 9, 2, 9}))
}
