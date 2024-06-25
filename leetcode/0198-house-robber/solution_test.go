package leetcode0198

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
			nums: []int{0},
			want: 0,
		},
		{
			name: "test 2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "test 3",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "test 4",
			nums: []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240},
			want: 4173,
		},
	}

	for _, test := range tests {
		have := robV0(test.nums)
		assert.Equalf(t, test.want, have, "%s: robV0 results incorrect", test.name)
	}

	for _, test := range tests {
		have := robV1(test.nums)
		assert.Equalf(t, test.want, have, "%s: robV1 results incorrect", test.name)
	}
}
