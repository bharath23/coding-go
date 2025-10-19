package leetcode0054

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])
	result := []int{}
	left := 0
	right := col - 1
	up := 0
	down := row - 1
	for len(result) < row*col {
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}

		for i := up + 1; i <= down; i++ {
			result = append(result, matrix[i][right])
		}

		if up != down {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[down][i])
			}
		}

		if left != right {
			for i := down - 1; i > up; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		up++
		down--
	}

	return result
}
