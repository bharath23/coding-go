package educative0028

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
			nums: []int{1, 2, 3, 5},
			want: 4,
		},
		{
			name: "test 2",
			nums: []int{2, 3, 4, 5, 6},
			want: 1,
		},
		{
			name: "test 3",
			nums: []int{-1, -2, -3, -4},
			want: 1,
		},
		{
			name: "test 4",
			nums: []int{55, 22, 52, 100, 1, 3, -5},
			want: 2,
		},
		{
			name: "test 5",
			nums: []int{0, -1, -2, -3, 5},
			want: 1,
		},
		{
			name: "test 6",
			nums: []int{1, 1},
			want: 2,
		},
	}

	for _, test := range tests {
		have := smallestMissingPositiveInteger(test.nums)
		assert.Equalf(t, test.want, have,
			"%s: smallestMissingPositiveInteger returned incorrect value",
			test.name)
	}
}
