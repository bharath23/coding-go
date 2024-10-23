package educative0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 6, 8, 10},
			target: 1,
			want:   0,
		}, {
			name:   "test 2",
			nums:   []int{11, 22, 33, 44, 55, 66, 77},
			target: 33,
			want:   2,
		}, {
			name:   "test 3",
			nums:   []int{-3, -1, 0, 11, 15},
			target: 0,
			want:   2,
		}, {
			name:   "test 4",
			nums:   []int{-30, -27, -8, -6, -1},
			target: -1,
			want:   4,
		}, {
			name:   "test 5",
			nums:   []int{0},
			target: 0,
			want:   0,
		},
	}

	for _, test := range tests {
		have := binarySearch(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s: incorrect binarySearch implementation", test.name)
	}
}
