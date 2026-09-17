package P36

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// fmt.Print(string(board[i][j]) + " ")
			if rowMap[i][board[i][j]] == true {
				return false
			}
			rowMap[i][board[i][j]] = true

			if lineMap[i][board[i][j]] == true {
				return false
			}
			lineMap[i][board[i][j]] = true

			if lineMap[i][board[i][j]] == true {
				return false
			}
			lineMap[i][board[i][j]] = true

		}
		fmt.Println()
	}
	return true
}
