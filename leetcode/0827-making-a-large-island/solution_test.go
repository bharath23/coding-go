package leetcode0827

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
			grid: [][]int{{1, 0}, {0, 1}},
			want: 3,
		},
		{
			name: "test 2",
			grid: [][]int{{1, 1}, {0, 1}},
			want: 4,
		},
		{
			name: "test 3",
			grid: [][]int{{1, 1}, {1, 1}},
			want: 4,
		},
	}

	for _, test := range tests {
		have := largestIsland(test.grid)
		assert.Equalf(t, test.want, have, "%s: largestIsland failed", test.name)
	}
}
