package leetcode0054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "test 1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "test 2",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	for _, test := range tests {
		have := spiralOrder(test.matrix)
		assert.Equalf(t, test.want, have, "%s: spiralOrder failed", test.name)
	}
}
