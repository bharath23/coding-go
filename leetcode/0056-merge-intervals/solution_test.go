package leetcode0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "test 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "test 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "test 3",
			intervals: [][]int{{1, 4}, {0, 1}},
			want:      [][]int{{0, 4}},
		},
	}

	for _, test := range tests {
		have := merge(test.intervals)
		assert.Equalf(t, test.want, have, "%s: merge failed", test.name)
	}
}
