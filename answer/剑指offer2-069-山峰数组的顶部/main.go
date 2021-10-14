package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	start, end := 0, len(arr)-1
	for start+1 < end {
		mid := (start + end) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid+1] {
			start = mid
		} else if arr[mid] > arr[mid+1] {
			end = mid
		}
	}
	return 0
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
}
