package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	record := map[string]int{}
	for i, v := range list1 {
		record[v] = i
	}

	min := math.MaxInt32
	result := make([]string, 0)
	for j, v := range list2 {
		if i, ok := record[v]; ok {
			if i+j < min {
				min = i + j
				result = []string{v}
			} else if i+j == min {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {

}
