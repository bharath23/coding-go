package leetcode0560

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
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "test 2",
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
		{
			name: "test 3",
			nums: []int{3, 4, 7, 2, -3, 1, 4, 2},
			k:    7,
			want: 4,
		},
	}

	for _, test := range tests {
		have := subarraySumv0(test.nums, test.k)
		assert.Equalf(t, test.want, have, "%s: subarraySumv0 failed", test.name)
		have = subarraySumv1(test.nums, test.k)
		assert.Equalf(t, test.want, have, "%s: subarraySumv1 failed", test.name)
	}
}
