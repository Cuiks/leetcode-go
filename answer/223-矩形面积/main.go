package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	if ax1 >= bx1 && ay1 >= by1 && ax2 <= bx2 && ay2 <= by2 {
		return (bx2 - bx1) * (by2 - by1)
	}
	if ax1 <= bx1 && ay1 <= by1 && ax2 >= bx2 && ay2 >= by2 {
		return (ax2 - ax1) * (ay2 - ay1)
	}
	area := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	if ax2 <= bx1 || bx2 <= ax1 || ay2 <= by1 || by2 <= ay1 {
		return area
	}
	x := min(abs(ax2-bx1), abs(bx2-ax1), abs(ax1-ax2), abs(bx1-bx2))
	y := min(abs(ay2-by1), abs(by2-ay1), abs(ay1-ay2), abs(by1-by2))
	fmt.Println(area, x, y, x*y)
	return area - x*y
}

func min(val ...int) int {
	res := val[0]
	for _, v := range val {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(computeArea(-2, -2, 2, 2, -3, -3, 3, -1))
}
