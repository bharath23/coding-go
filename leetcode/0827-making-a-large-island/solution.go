package leetcode0827

/*
 * The solution consists of two main phases:
 * 1. A DFS pass that visits each cell at most once to label islands and
 * 	  compute their sizes.
 * 2. A scanning pass over all cells; for each 0-cell, we check at most 4
 *    neighbours (constant time) to compute a potential size if flipped.
 *
 * Each of these passes is O(n^2) for an n×n grid, so the overall time
 * complexity is O(n^2).
 *
 * Additional space is used for:
 *  - the `islandId` 2-D array of size n×n,
 *  - the recursion stack in DFS which in the worst case can be O(n^2),
 *  - the `sizes` map which in the worst case can have O(n^2) entries.
 * Hence the space complexity is O(n^2).
 *
 * Time Complexity:  O(n^2)
 * Space Complexity: O(n^2)
 */
func largestIsland(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	islandId := make([][]int, n)
	for i := 0; i < n; i++ {
		islandId[i] = make([]int, n)
	}

	islandSize := make(map[int]int)
	curId := 2

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= n {
			return 0
		}

		if grid[i][j] != 1 || islandId[i][j] != 0 {
			return 0
		}

		islandId[i][j] = curId
		size := 1
		for _, d := range directions {
			size += dfs(i+d[0], j+d[1])
		}

		return size
	}

	maxSize := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && islandId[i][j] == 0 {
				size := dfs(i, j)
				islandSize[curId] = size
				curId++
				if maxSize < size {
					maxSize = size
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				islandSeen := map[int]bool{}
				size := 1
				for _, d := range directions {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < n {
						id := islandId[ni][nj]
						if id > 1 && !islandSeen[id] {
							size += islandSize[id]
							islandSeen[id] = true
						}
					}
				}
				if size > maxSize {
					maxSize = size
				}
			}
		}
	}

	return maxSize
}
