package leetcode0034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "test 1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "test 2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "test 1",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, test := range tests {
		have := searchRange(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s: searchRange failed", test.name)
	}
}
