package array

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

func IsValidSudoku(board [][]byte) bool {

	return rowsValid(board) && colsValid(board) && allCubesValid(board)

}

func allCubesValid(board [][]byte) bool {

	for row := 0; row < 9; row = row + 3 {

		for col := 0; col < 9; col = col + 3 {

			if !isValidCube(board, row, col) {
				return false
			}

		}

	}

	return true
}

func colsValid(board [][]byte) bool {

	for col := 0; col < 9; col++ {
		if !isValidCol(board, col) {
			return false
		}
	}

	return true
}

func rowsValid(board [][]byte) bool {

	for row := 0; row < 9; row++ {
		if !isValidRow(board, row) {
			return false
		}
	}

	return true
}

func isValidCube(board [][]byte, row, col int) bool {

	m := make(map[byte]struct{})

	for r := row; r < (row + 3); r++ {

		for c := col; c < (col + 3); c++ {

			if board[r][c] == '.' {
				continue
			}

			if _, ok := m[board[r][c]]; ok {
				return false
			}

			m[board[r][c]] = struct{}{}
		}

	}

	return true
}

func isValidCol(board [][]byte, col int) bool {

	m := make(map[byte]struct{})

	for row := 0; row < 9; row++ {

		if board[row][col] == '.' {
			continue
		}

		if _, ok := m[board[row][col]]; ok {
			return false
		}

		m[board[row][col]] = struct{}{}

	}

	return true
}

func isValidRow(board [][]byte, row int) bool {

	m := make(map[byte]struct{})

	for col := 0; col < 9; col++ {

		if board[row][col] == '.' {
			continue
		}

		if _, ok := m[board[row][col]]; ok {
			return false
		}

		m[board[row][col]] = struct{}{}

	}

	return true
}
