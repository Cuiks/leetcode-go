package main

//func findNumberIn2DArray(matrix [][]int, target int) bool {
//	if len(matrix) == 0 || len(matrix[0]) == 0 {
//		return false
//	}
//	for i := 0; i < len(matrix); i++ {
//		if binary(matrix[i], target) {
//			return true
//		}
//	}
//	return false
//}
//
//func binary(list []int, target int) bool {
//	start := 0
//	end := len(list) - 1
//	for start+1 < end {
//		mid := start + (end-start)/2
//		if list[mid] == target {
//			return true
//		} else if target < list[mid] {
//			end = mid
//		} else {
//			start = mid
//		}
//	}
//	if list[start] == target || list[end] == target {
//		return true
//	}
//	return false
//}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if target == matrix[i][j] {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

func main() {

}
