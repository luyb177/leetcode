package main

// func main() {}

func solveSudoku(board [][]byte) {
	backTracking(board, 0, 0)
}

func backTracking(board [][]byte, row, col int) bool {
	// 先处理行切换
	if col == 9 {
		// 切换到下一行，列重置为0
		return backTracking(board, row+1, 0)
	}

	// 终止条件：所有行都处理完
	if row == 9 {
		return true
	}

	// 如果当前位置已经有数字，直接处理下一个位置
	if board[row][col] != '.' {
		return backTracking(board, row, col+1)
	}

	// 当前位置是空的，尝试1-9
	for v := '1'; v <= '9'; v++ { // 修正：包含'9'
		if isValid(board, row, col, byte(v)) {
			board[row][col] = byte(v)
			if backTracking(board, row, col+1) {
				return true
			}
			board[row][col] = '.' // 回溯
		}
	}
	// 9个数字没有递归进去，回溯
	return false
}

func isValid(board [][]byte, row, col int, v byte) bool {
	// 行
	for j := 0; j < 9; j++ {
		if board[row][j] == v {
			return false
		}
	}

	// 列
	for i := 0; i < 9; i++ {
		if board[i][col] == v {
			return false
		}
	}

	// 九宫格
	var startRow = row / 3 * 3
	var startCol = col / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == v {
				return false
			}
		}
	}

	return true
}
