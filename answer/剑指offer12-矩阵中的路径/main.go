package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := map[string]bool{}
	for i, row := range board {
		for j := range row {
			if backtrack(board, i, j, 0, word, visited) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, i, j, k int, word string, visited map[string]bool) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}
	if v, ok := visited[fmt.Sprintf("%d:%d", i, j)]; ok && v {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	visited[fmt.Sprintf("%d:%d", i, j)] = true
	if backtrack(board, i+1, j, k+1, word, visited) || backtrack(board, i-1, j, k+1, word, visited) || backtrack(board, i, j+1, k+1, word, visited) || backtrack(board, i, j-1, k+1, word, visited) {
		return true
	}
	visited[fmt.Sprintf("%d:%d", i, j)] = false
	return false
}

func main() {
	board := [][]byte{{'a', 'a'}}
	word := "aaa"
	fmt.Println(exist(board, word))
}
