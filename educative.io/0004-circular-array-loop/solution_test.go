package educative0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "test1",
			nums: []int{1, 3, -2, -4, 1},
			want: true,
		},
		{
			name: "test2",
			nums: []int{2, 1, -2, -2},
			want: false,
		},
		{
			name: "test3",
			nums: []int{5, 4, -2, -1, 3},
			want: false,
		},
		{
			name: "test4",
			nums: []int{1, 2, -3, 3, 4, 7},
			want: true,
		},
		{
			name: "test5",
			nums: []int{3, 3, 1, -1, 2},
			want: true,
		},
		{
			name: "test6",
			nums: []int{-4737,-4455},
			want: true,
		},
	}

	for _, test := range tests {
		have := circularArrayLoop(test.nums)
		assert.Equalf(t, test.want, have, "%s: circularArrayLoop incorrect result", test.name)
	}
}
