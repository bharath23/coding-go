package educative0012

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
			nums:   []int{6, 7, 1, 2, 3, 4, 5},
			target: 3,
			want:   4,
		}, {
			name:   "test 2",
			nums:   []int{6, 7, 1, 2, 3, 4, 5},
			target: 6,
			want:   0,
		}, {
			name:   "test 3",
			nums:   []int{4, 5, 6, 1, 2, 3},
			target: 3,
			want:   5,
		}, {
			name:   "test 4",
			nums:   []int{4, 5, 6, 1, 2, 3},
			target: 6,
			want:   2,
		}, {
			name:   "test 5",
			nums:   []int{4},
			target: 1,
			want:   -1,
		},
	}

	for _, test := range tests {
		have := binarySearchRotated(test.nums, test.target)
		assert.Equalf(t, test.want, have, "%s: incorrect binarySearch implementation", test.name)
	}
}
