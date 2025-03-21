package educative0021

func setMatrixZeros(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	fcol := false
	frow := false
	for i := 0; i < m; i++ {
		if mat[i][0] == 0 {
			fcol = true
		}
	}

	for i := 0; i < n; i++ {
		if mat[0][i] == 0 {
			frow = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if mat[i][0] == 0 {
			for j := 1; j < n; j++ {
				mat[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if mat[0][j] == 0 {
			for i := 1; i < m; i++ {
				mat[i][j] = 0
			}
		}
	}

	if fcol {
		for i := 0; i < m; i++ {
			mat[i][0] = 0
		}
	}

	if frow {
		for j := 0; j < n; j++ {
			mat[0][j] = 0
		}
	}

	return mat
}
