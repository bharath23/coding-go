package leetcode1091

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "test 1",
			grid: [][]int{[]int{0, 1}, []int{1, 0}},
			want: 2,
		},
		{
			name: "test 2",
			grid: [][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}},
			want: 4,
		},
		{
			name: "test 3",
			grid: [][]int{[]int{1, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}},
			want: -1,
		},
	}

	for _, test := range tests {
		have := shortestPathBinaryMatrix(test.grid)
		assert.Equalf(t, test.want, have, "%s: shortest path length mismatch", test.name)
	}
}

func TestSolutionFB(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want string
	}{
		{
			name: "test 1",
			grid: [][]int{[]int{0, 1}, []int{1, 0}},
			want: "",
		},
		{
			name: "test 2",
			grid: [][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}},
			want: "(0, 0) -> (0, 1) -> (0, 2) -> (1, 2) -> (2, 2)",
		},
		{
			name: "test 3",
			grid: [][]int{[]int{1, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}},
			want: "",
		},
		{
			name: "test 4",
			grid: [][]int{
				[]int{0, 0, 0, 0, 0, 0, 0},
				[]int{0, 0, 1, 0, 0, 1, 0},
				[]int{0, 0, 1, 0, 1, 1, 0},
				[]int{0, 0, 1, 0, 1, 0, 1},
				[]int{1, 1, 1, 0, 0, 0, 0},
			},
			want: "(0, 0) -> (0, 1) -> (0, 2) -> (0, 3) -> (1, 3) -> (2, 3) -> (3, 3) -> (4, 3) -> (4, 4) -> (4, 5) -> (4, 6)",
		},
	}

	for _, test := range tests {
		have := shortestPathBinaryMatrixFB(test.grid)
		assert.Equalf(t, test.want, have, "%s: shortest path length mismatch", test.name)
	}
}
