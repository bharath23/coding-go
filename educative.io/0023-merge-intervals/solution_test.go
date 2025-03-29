package educative0023

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
			intervals: [][]int{{1, 5}, {3, 7}, {4, 6}},
			want:      [][]int{{1, 7}},
		},
		{
			name:      "test 2",
			intervals: [][]int{{1, 5}, {4, 6}, {6, 8}, {11, 15}},
			want:      [][]int{{1, 8}, {11, 15}},
		},
		{
			name:      "test 3",
			intervals: [][]int{{1, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "test 4",
			intervals: [][]int{{1, 9}, {3, 8}, {4, 4}},
			want:      [][]int{{1, 9}},
		},
		{
			name:      "test 5",
			intervals: [][]int{{1, 2}, {3, 4}, {8, 8}},
			want:      [][]int{{1, 2}, {3, 4}, {8, 8}},
		},
	}

	for _, test := range tests {
		have := mergeIntervals(test.intervals)
		assert.Equalf(t, test.want, have, "%s: mergeIntervals returned "+
			"incorrect intervals", test.name)

	}
}
