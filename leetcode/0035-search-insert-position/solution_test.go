package leetcode0035

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
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "test 1",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	for _, test := range tests {
		have := searchInsert(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s: searchInsert failed", test.name)
	}
}
