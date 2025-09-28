package leetcode0073

func setZeroesv0(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesv1(matrix [][]int) {
	isCol := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if isCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
