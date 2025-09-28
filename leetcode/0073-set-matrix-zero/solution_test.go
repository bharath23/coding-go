package leetcode0073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "test 1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "test 2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name:   "test 2",
			matrix: [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}},
			want:   [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
		},
	}

	for _, test := range tests {
		matrix := make([][]int, len(test.matrix))
		for i := range matrix {
			matrix[i] = make([]int, len(test.matrix[0]))
			copy(matrix[i], test.matrix[i])
		}
		setZeroesv0(matrix)
		assert.Equalf(t, test.want, matrix, "%s: setZeroesv0 failed", test.name)

		for i := range matrix {
			matrix[i] = make([]int, len(test.matrix[0]))
			copy(matrix[i], test.matrix[i])
		}
		setZeroesv1(matrix)
		assert.Equalf(t, test.want, matrix, "%s: setZeroesv1 failed", test.name)
	}
}
