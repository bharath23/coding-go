package leetcode1091

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	type element struct {
		i, j    int
		pathLen int
	}

	queue := []element{}
	enqueue := func(e element) {
		queue = append(queue, e)
	}

	dequeue := func() element {
		e := queue[0]
		queue = queue[1:]
		return e
	}

	isQueueEmpty := func() bool {
		return (len(queue) == 0)
	}

	n := len(grid)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	findNeighbors := func(l, m, length int) []element {
		rows := []int{-1, -1, -1, 0, 0, 1, 1, 1}
		cols := []int{-1, 0, 1, -1, 1, -1, 0, 1}
		ee := []element{}
		for k := 0; k < len(rows); k++ {
			i := l + rows[k]
			j := m + cols[k]
			if (i >= 0) && (i < n) && (j >= 0) && (j < n) &&
				(grid[i][j] == 0) &&
				!visited[i][j] {
				ee = append(ee, element{i, j, length + 1})
			}
		}

		return ee
	}

	if (grid[0][0] != 0) || (grid[n-1][n-1] != 0) {
		return -1
	}

	enqueue(element{0, 0, 1})
	visited[0][0] = true
	for !isQueueEmpty() {
		e := dequeue()
		if (e.i == n-1) && (e.j == n-1) {
			return e.pathLen
		}

		for _, n := range findNeighbors(e.i, e.j, e.pathLen) {
			visited[n.i][n.j] = true
			enqueue(n)
		}
	}

	return -1
}

type element struct {
	i, j int
	path []element
}

func shortestPathBinaryMatrixFB(grid [][]int) string {
	queue := []element{}
	enqueue := func(e element) {
		queue = append(queue, e)
	}

	dequeue := func() element {
		e := queue[0]
		queue = queue[1:]
		return e
	}

	isQueueEmpty := func() bool {
		return (len(queue) == 0)
	}

	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	findNeighbors := func(p, q int, path []element) []element {
		rows := []int{-1, 0, 0, 1}
		cols := []int{0, -1, 1, 0}
		ee := []element{}
		for k := 0; k < len(rows); k++ {
			i := p + rows[k]
			j := q + cols[k]
			if (i >= 0) && (i < m) && (j >= 0) && (j < n) &&
				(grid[i][j] == 0) &&
				!visited[i][j] {
				pp := make([]element, len(path))
				copy(pp, path)
				ee = append(ee, element{i, j, append(pp, element{i, j, nil})})
			}
		}

		return ee
	}

	pathToString := func(path []element) string {
		if len(path) == 0 {
			return ""
		}

		str := fmt.Sprintf("(%d, %d)", path[0].i, path[0].j)
		for i := 1; i < len(path); i++ {
			str += fmt.Sprintf(" -> (%d, %d)", path[i].i, path[i].j)
		}

		return str
	}

	if (grid[0][0] != 0) || (grid[m-1][n-1] != 0) {
		return pathToString(nil)
	}

	enqueue(element{0, 0, []element{element{0, 0, nil}}})
	visited[0][0] = true
	for !isQueueEmpty() {
		e := dequeue()
		if (e.i == m-1) && (e.j == n-1) {
			return pathToString(e.path)
		}

		for _, n := range findNeighbors(e.i, e.j, e.path) {
			visited[n.i][n.j] = true
			enqueue(n)
		}
	}

	return pathToString(nil)
}
