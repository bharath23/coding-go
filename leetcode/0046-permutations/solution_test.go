package leetcod0046

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "test 2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "test 1",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, test := range tests {
		have := permutev0(test.nums)
		assert.ElementsMatchf(t, test.want, have, "%s: permute failed",
			test.name)
	}

	for _, test := range tests {
		have := permutev1(test.nums)
		assert.ElementsMatchf(t, test.want, have, "%s: permute failed",
			test.name)
	}
}
