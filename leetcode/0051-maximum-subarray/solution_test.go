package leetcode0051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "test 2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "test 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}

	for _, test := range tests {
		have := maxSubArrayv0(test.nums)
		assert.Equalf(t, test.want, have, "%s: maxSubArrayv0 failed", test.name)
	}

	for _, test := range tests {
		have := maxSubArrayv1(test.nums)
		assert.Equalf(t, test.want, have, "%s: maxSubArrayv1 failed", test.name)
	}
}
