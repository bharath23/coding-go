package meta0003

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
			nums: []int{1, 2, 3, 4, 5, 5},
			want: 4,
		},
		{
			name: "test 2",
			nums: []int{4, 1, 2, 3},
			want: 2,
		},
		{
			name: "test 3",
			nums: []int{4, 3, 2, 1},
			want: -1,
		},
		{
			name: "test 4",
			nums: []int{1, 2, 1, 1, 3},
			want: 3,
		},
		{
			name: "test 5",
			nums: []int{1, 1, 1, 1, 1, 5},
			want: 5,
		},
		{
			name: "test 6",
			nums: []int{5, 2, 3},
			want: 1,
		},
	}

	for _, test := range tests {
		have := canSplitArray(test.nums)
		assert.Equalf(t, test.want, have, "%s: canSplitArray return mismatch",
			test.name)
	}
}
