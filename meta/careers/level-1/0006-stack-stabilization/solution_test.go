package meta0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		r    []int
		want int
	}{
		{
			name: "test 1",
			n:    5,
			r:    []int{2, 5, 3, 6, 5},
			want: 3,
		},
		{
			name: "test 2",
			n:    3,
			r:    []int{100, 100, 100},
			want: 2,
		},
		{
			name: "test 1",
			n:    4,
			r:    []int{6, 5, 4, 3},
			want: -1,
		},
	}

	for _, test := range tests {
		have := getMinimumDeflatedDiscCount(test.n, test.r)
		assert.Equalf(t, test.want, have, "%s: getMinimumDeflatedDiscCount failed",
			test.name)
	}
}
