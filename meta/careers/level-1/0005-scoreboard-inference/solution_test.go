package meta0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		s    []int
		want int
	}{
		{
			name: "test 1",
			n:    6,
			s:    []int{1, 2, 3, 4, 5, 6},
			want: 4,
		},
		{
			name: "test 2",
			n:    4,
			s:    []int{4, 3, 3, 4},
			want: 3,
		},
		{
			name: "test 3",
			n:    4,
			s:    []int{2, 4, 6, 8},
			want: 4,
		},
	}

	for _, test := range tests {
		have := getMinProblemCount(test.n, test.s)
		assert.Equalf(t, test.want, have, "%s, getMinProblemCount failed",
			test.name)
	}
}
