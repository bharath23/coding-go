package leetcode0547

import "github.com/bharath23/coding-go/internal"

/*
A simple DFS solution. We visit each node and traverse its children if we
have not already visited it. Total time complexity for this O(n^2) as we
traverse the entire adjacency matrix which is n x n. Space complexity is
O(n) as we store a visited map.

Complexity:
Time complexity: O(n^2)
Space complexity: O(n)
*/

func findCircleNumV0(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	var dfs func(int)
	dfs = func(node int) {
		for i := 0; i < len(isConnected[node]); i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				visited[i] = true
				dfs(i)
			}
		}
	}

	count := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}

/*
A simple BFS solution. We visit each node and queue its children if we
have not already visited it. Total time complexity for this O(n^2) as we
traverse the entire adjacency matrix which is n x n. Space complexity is
O(n) as we store a visited map.

Complexity:
Time complexity: O(n^2)
Space complexity: O(n)
*/

func findCircleNumV1(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	count := 0
	queue := internal.NewQueue()
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			queue.Enqueue(i)
			for !queue.IsEmpty() {
				n, _ := queue.Dequeue()
				visited[n] = true
				for j := 0; j < len(isConnected[n]); j++ {
					if isConnected[n][j] == 1 && !visited[j] {
						queue.Enqueue(j)
					}
				}
			}
			count++
		}
	}

	return count
}

/*
Solution based on union-find. Time complexity for union is O(n). Time
complexity to create parents array is O(n^3). Space complexity is O(n) for
storing parents array.

Complexity:
Time complexity: O(n^3)
Space complexity: O(n)
*/

func findCircleNumV2(isConnected [][]int) int {
	parents := make([]int, len(isConnected))
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}

	var find func(int) int
	find = func(i int) int {
		if parents[i] == -1 {
			return i
		}

		return find(parents[i])
	}

	union := func(i, j int) {
		parentI := find(i)
		parentJ := find(j)
		if parentI != parentJ {
			parents[parentI] = parentJ
		}
	}

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 && i != j {
				union(i, j)
			}
		}
	}

	count := 0
	for i := 0; i < len(isConnected); i++ {
		if parents[i] == -1 {
			count++
		}
	}

	return count
}
