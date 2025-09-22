package leetcode0051

/*
 * The algorithm uses backtracking to explore all possible placements of queens
 * on an n x n chessboard. For each row, it tries every column and places a
 * queen if the column, diagonal, and anti-diagonal are not already attacked.
 * Valid placements are recorded, and the algorithm backtracks to try
 * alternative configurations.
 *
 * In the worst case, the backtracking tree explores O(n!) possibilities since
 * each row places one queen in a unique column. At each recursive call, checks
 * for conflicts in sets (columns, diagonals, anti-diagonals) are performed in
 * constant time, so the dominant factor is the exponential search space.
 * Therefore, the overall time complexity is O(n!).
 *
 * The space complexity is O(n^2) to store the board configuration and O(n) for
 * recursion depth, with additional O(n) storage in hash sets for columns and
 * diagonals. Thus the space complexity is O(n^2).
 *
 * Time complexity: O(n!)
 * Space complexity: O(n^2)
 */
func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	hasQueen := func(set map[int]bool, pos int) bool {
		_, ok := set[pos]
		return ok
	}

	addQueen := func(set map[int]bool, pos int) {
		set[pos] = true
	}

	removeQueen := func(set map[int]bool, pos int) {
		delete(set, pos)
	}

	createBoard := func(board [][]rune) []string {
		strBoard := make([]string, n)
		for row := range board {
			strBoard[row] = string(board[row])
		}

		return strBoard
	}

	var backtrack func(row int, columns, diagonals, antiDiagonals map[int]bool, board [][]rune)

	backtrack = func(row int, columns, diagonals, antiDiagonals map[int]bool, board [][]rune) {
		if row == n {
			result = append(result, createBoard(board))
		}

		for col := 0; col < n; col++ {
			diagonal := row - col
			antiDiagonal := row + col

			// Can queen be placed in this column
			if hasQueen(columns, col) ||
				hasQueen(diagonals, diagonal) ||
				hasQueen(antiDiagonals, antiDiagonal) {
				continue
			}

			addQueen(columns, col)
			addQueen(diagonals, diagonal)
			addQueen(antiDiagonals, antiDiagonal)
			board[row][col] = 'Q'

			// Find the position of queen in the next row
			backtrack(row+1, columns, diagonals, antiDiagonals, board)

			// Reset the board before moving to the next possibility
			board[row][col] = '.'
			removeQueen(antiDiagonals, antiDiagonal)
			removeQueen(diagonals, diagonal)
			removeQueen(columns, col)
		}
	}

	backtrack(0, make(map[int]bool), make(map[int]bool), make(map[int]bool), board)

	return result
}
