package meta0004

import (
	"fmt"
	"math/rand"
)

func generateMinesOnGrid(m, n, k int) []string {
	// This approach utilizes the reservoir sampling algorithm. Reservoir
	// sampling is a randomized algorithm used to select a random sample of k
	// items from a population of unknown size in a single pass. If you’re
	// unfamiliar with this algorithm, I recommend reading about it. It’s the
	// only way to achieve O(n) time complexity for this task.
	generateKMines := func(n, k int) []int {
		res := make([]int, k)
		for i := range k {
			res[i] = i
		}

		for i := k; i < n; i++ {
			j := rand.Intn(i)
			if j < k {
				res[j] = i
			}
		}

		return res
	}

	mines := []string{}
	loc := generateKMines(m*n, k)
	for _, l := range loc {
		i := l / m
		j := l % m
		mines = append(mines, fmt.Sprintf("(%d, %d)", i, j))
	}

	return mines
}
