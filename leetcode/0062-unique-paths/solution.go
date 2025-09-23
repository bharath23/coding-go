package leetcode0062

/*
 * The algorithm uses dynamic programming to fill an m x n grid where each
 * cell d[i][j] represents the number of unique paths to reach that cell.
 * The first row and first column are initialized to 1, since there is only
 * one way to reach those cells (by moving only right or only down). For
 * every other cell, the number of paths is the sum of the cell above and
 * the cell to the left. The final result is stored in d[m-1][n-1].
 *
 * The DP table requires filling m*n cells, and each cell computation takes
 * O(1) time, so the total time complexity is O(m*n).
 *
 * The space complexity is O(m*n) for storing the DP table.
 *
 * Time complexity: O(m*n)
 * Space complexity: O(m*n)
 */
func uniquePaths(m, n int) int {
	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				d[i][j] = 1
			} else {
				d[i][j] = d[i-1][j] + d[i][j-1]
			}
		}
	}

	return d[m-1][n-1]
}
