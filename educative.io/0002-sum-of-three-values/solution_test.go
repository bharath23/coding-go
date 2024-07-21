package educative0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "test1",
			nums:   []int{1, -1, 0},
			target: -1,
			want:   false,
		},
		{
			name:   "test2",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 10,
			want:   true,
		},
		{
			name:   "test3",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 21,
			want:   false,
		},
		{
			name:   "test1",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			want:   true,
		},
		{
			name:   "test5",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			want:   true,
		},
	}

	for _, test := range tests {
		dst := make([]int, len(test.nums))
		copy(dst, test.nums)
		have := findSumOfThree(dst, test.target)
		assert.Equalf(t, test.want, have, "%s: findSumOfThree incorrect result", test.name)
	}
}
