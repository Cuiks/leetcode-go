package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			result[j-1] += bookings[i][2]
		}
	}
	return result
}

func main() {
	fmt.Println(corpFlightBookings([][]int{{1,2,10},{2,3,20},{2,5,25}}, 5))
}
