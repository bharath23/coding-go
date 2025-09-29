package leetcode0078

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
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "test 1",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}

	for _, test := range tests {
		have := subsetsv0(test.nums)
		assert.ElementsMatchf(t, test.want, have, "%s: subsetsv0 failed", test.name)
		have = subsetsv1(test.nums)
		assert.ElementsMatchf(t, test.want, have, "%s: subsetsv1 failed", test.name)
		have = subsetsv2(test.nums)
		assert.ElementsMatchf(t, test.want, have, "%s: subsetsv2 failed", test.name)
	}
}
