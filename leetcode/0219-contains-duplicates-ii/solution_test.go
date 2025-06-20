package leetcode0219

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "test 2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "test 3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}

	for _, test := range tests {
		have := containsNearbyDuplicatev0(test.nums, test.k)
		assert.Equalf(t, test.want, have, "%s: containsNearbyDuplicatv0 failed",
			test.name)
	}

	for _, test := range tests {
		have := containsNearbyDuplicatev1(test.nums, test.k)
		assert.Equalf(t, test.want, have, "%s: containsNearbyDuplicatv1 failed",
			test.name)
	}
}
