package leetcode0740

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
			nums: []int{3, 4, 2},
			want: 6,
		},
		{
			name: "test 2",
			nums: []int{2, 2, 3, 3, 3, 4},
			want: 9,
		},
	}

	for _, test := range tests {
		have := deleteAndEarnV0(test.nums)
		assert.Equalf(t, test.want, have, "%s: deleteAndEarnV0 result mismatch", test.name)
	}

	for _, test := range tests {
		have := deleteAndEarnV1(test.nums)
		assert.Equalf(t, test.want, have, "%s: deleteAndEarnV1 result mismatch", test.name)
	}

	for _, test := range tests {
		have := deleteAndEarnV2(test.nums)
		assert.Equalf(t, test.want, have, "%s: deleteAndEarnV2 result mismatch", test.name)
	}
}
