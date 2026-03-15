package main

//func main() {}
var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	backtracking(n, 0, board)

	return res
}

func backtracking(n int, row int, board [][]byte) {
	if row == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = string(board[i])
		}
		res = append(res, temp)
		return
	}

	for col := 0; col < n; col++ {
		if !isValid(row, col, board, n) {
			continue
		}

		board[row][col] = 'Q'
		backtracking(n, row+1, board)
		board[row][col] = '.'
	}
}

func isValid(row, col int, board [][]byte, n int) bool {
	// 检查列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 右上
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
