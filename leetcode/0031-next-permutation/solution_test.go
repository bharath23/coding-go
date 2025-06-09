package leetcode0031

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "test 2",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "test 3",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
	}

	for _, test := range tests {
		nextPermutation(test.nums)
		assert.Equalf(t, test.want, test.nums, "%s: nextPermutation failed",
			test.name)
	}
}
