package educative0022

func rotateImage(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n-1; i++ {
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
	return matrix
}
