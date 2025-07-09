package leetcode0215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "test 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "test 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := findKthLargestv0(nums, test.k)
		assert.Equalf(t, test.want, have, "%s: findKthLargestv0 failed", test.name)
	}

	for _, test := range tests {
		have := findKthLargestv1(test.nums, test.k)
		assert.Equalf(t, test.want, have, "%s: findKthLargestv1 failed", test.name)
	}

	for _, test := range tests {
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)
		have := findKthLargestv2(nums, test.k)
		assert.Equalf(t, test.want, have, "%s: findKthLargestv0 failed", test.name)
	}
}
