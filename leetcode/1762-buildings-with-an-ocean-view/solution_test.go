package leetcode1762

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    []int
	}{
		{
			name:    "test 1",
			heights: []int{4, 2, 3, 1},
			want:    []int{0, 2, 3},
		},
		{
			name:    "test 2",
			heights: []int{4, 3, 2, 1},
			want:    []int{0, 1, 2, 3},
		},
		{
			name:    "test 2",
			heights: []int{1, 3, 2, 4},
			want:    []int{3},
		},
		{
			name:    "test 3",
			heights: []int{4, 5, 4, 2, 3, 1, 1},
			want:    []int{1, 2, 4, 6},
		},
	}

	for _, test := range tests {
		have := findBuildings(test.heights)
		assert.Equalf(t, test.want, have, "%s: findBuildings failed", test.name)
	}
}
