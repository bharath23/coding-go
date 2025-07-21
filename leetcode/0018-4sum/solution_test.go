package leetcode0018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:   "test 2",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := fourSumv0(nums, test.target)
		assert.ElementsMatchf(t, test.want, have, "%s, fourSumv0 failed",
			test.name)
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := fourSumv1(nums, test.target)
		assert.ElementsMatchf(t, test.want, have, "%s, fourSumv1 failed",
			test.name)
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := fourSum(nums, test.target)
		assert.ElementsMatchf(t, test.want, have, "%s, fourSum failed",
			test.name)
	}
}
