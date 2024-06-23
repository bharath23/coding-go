package leetcode0261

import (
	"github.com/bharath23/coding-go/internal"
)

func validTree(n int, edges [][]int) bool {
	adjMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		adjMatrix[i] = make([]int, n)
	}

	for _, e := range edges {
		adjMatrix[e[0]][e[1]] = 1
		adjMatrix[e[1]][e[0]] = 1
	}

	parents := make(map[int]int)
	stack := internal.NewStack()
	stack.Push(0)
	parents[0] = -1
	for !stack.IsEmpty() {
		node := stack.Pop().(int)
		for i := 0; i < n; i++ {
			if adjMatrix[node][i] == 0 {
				continue
			}

			if parents[node] == i {
				continue
			}

			_, ok := parents[i]
			if ok {
				return false
			}

			stack.Push(i)
			parents[i] = node
		}
	}

	return len(parents) == n
}
