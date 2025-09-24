package leetcode0066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name:   "test 1",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "test 2",
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			name:   "test 3",
			digits: []int{9},
			want:   []int{1, 0},
		},
	}

	for _, test := range tests {
		have := plusOne(test.digits)
		assert.Equalf(t, test.want, have, "%s: plusOne failed", test.name)
	}
}
