package educative0022

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
			matrix: [][]int{{6, 9}, {2, 7}},
			want:   [][]int{{2, 6}, {7, 9}},
		},
		{
			name:   "test 2",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name:   "test 3",
			matrix: [][]int{{2, 14, 8}, {12, 7, 14}, {3, 3, 7}},
			want:   [][]int{{3, 12, 2}, {3, 7, 14}, {7, 14, 8}},
		},
		{
			name: "test 4",
			matrix: [][]int{{3, 1, 1, 7}, {15, 12, 13, 13}, {4, 14, 12, 4},
				{10, 5, 11, 12}},
			want: [][]int{{10, 4, 15, 3}, {5, 14, 12, 1}, {11, 12, 13, 1},
				{12, 4, 13, 7}},
		},
		{
			name: "test 5",
			matrix: [][]int{{10, 1, 14, 11, 14}, {13, 4, 8, 2, 13},
				{10, 19, 1, 6, 8}, {20, 10, 8, 2, 12}, {15, 6, 8, 8, 18}},
			want: [][]int{{15, 20, 10, 13, 10}, {6, 10, 19, 4, 1},
				{8, 8, 1, 8, 14}, {8, 2, 6, 2, 11}, {18, 12, 8, 13, 14}},
		},
	}

	for _, test := range tests {
		have := rotateImage(test.matrix)
		assert.Equalf(t, test.want, have, "%s: rotateImage incorrect output",
			test.name)
	}
}
