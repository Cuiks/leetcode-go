package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		record := map[byte]bool{}
		for j := 0; j < len(board[0]); j++ {
			ch := board[i][j]
			if ch == '.'{
				continue
			}
			if record[ch] {
				return false
			} else {
				record[ch] = true
			}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		record := map[byte]bool{}
		for j := 0; j < len(board); j++ {
			ch := board[j][i]
			if ch == '.'{
				continue
			}
			if record[ch] {
				return false
			} else {
				record[ch] = true
			}
		}
	}

	direction := [][]int{{0, 0}, {-1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, 0}, {-1, 1}, {1, -1}}
	for _, item := range [][]int{{1, 1}, {1, 4}, {1, 7}, {4, 1}, {4, 4}, {4, 7}, {7, 1}, {7, 4}, {7, 7}} {
		record := map[byte]bool{}
		for _, dire := range direction {
			i, j := item[0]+dire[0], item[1]+dire[1]
			ch := board[i][j]
			if ch == '.'{
				continue
			}
			if record[ch] {
				return false
			} else {
				record[ch] = true
			}
		}
	}
	return true
}

func main() {

}
