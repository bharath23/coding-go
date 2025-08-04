package leetcode0033

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "test 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "test 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "test 1",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
	}

	for _, test := range tests {
		have := search(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s: search failed", test.name)
	}
}
