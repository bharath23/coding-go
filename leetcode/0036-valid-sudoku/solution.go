package leetcode0036

/*
 * We iterate through all N × N cells of the Sudoku board. For each cell,
 * we perform constant-time lookups and insertions into hash maps for rows,
 * columns, and sub-boxes. Thus, each cell is processed in O(1) time.
 * Overall time complexity is O(N^2).

 * The space complexity is O(N^2) because we maintain up to N entries in
 * each of the N row, column, and sub-box maps.
 *
 * Time complexity: O(N^2)
 * Space complexity: O(N^2)
 */

func isValidSudokuv0(board [][]byte) bool {
	N := 9
	cols := make([]map[byte]bool, N)
	rows := make([]map[byte]bool, N)
	subBoxes := make([]map[byte]bool, N)
	for i := 0; i < N; i++ {
		cols[i] = make(map[byte]bool, N)
		rows[i] = make(map[byte]bool, N)
		subBoxes[i] = make(map[byte]bool, N)
	}

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}

			if cols[c][val] {
				return false
			}

			cols[c][val] = true
			if rows[r][val] {
				return false
			}

			rows[r][val] = true
			box := (r/3)*3 + c/3
			if subBoxes[box][val] {
				return false
			}

			subBoxes[box][val] = true
		}
	}

	return true
}

/*
 * We iterate through all N × N cells of the Sudoku board. For each cell,
 * we perform constant-time bitwise operations to check and update the state
 * of rows, columns, and sub-boxes. Thus, each cell is processed in O(1) time.
 * Overall time complexity is O(N^2).
 * The space complexity is O(N) since we only store N integers each for rows,
 * columns, and sub-boxes instead of hash maps.
 *
 * Time complexity: O(N^2)
 * Space complexity: O(N)
 */

func isValidSudokuv1(board [][]byte) bool {
	N := 9
	cols := make([]int, N)
	rows := make([]int, N)
	subBoxes := make([]int, N)

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if board[r][c] == '.' {
				continue
			}

			val := board[r][c] - '0'
			bit := 1 << (val - 1)
			if cols[c]&bit != 0 {
				return false
			}

			cols[c] |= bit
			if rows[r]&bit != 0 {
				return false
			}

			rows[r] |= bit
			box := (r/3)*3 + c/3
			if subBoxes[box]&bit != 0 {
				return false
			}

			subBoxes[box] |= bit
		}
	}

	return true
}
