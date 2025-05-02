package meta0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		r    int32
		c    int32
		g    [][]int32
		want float64
	}{
		{
			name: "test 1",
			r:    2,
			c:    3,
			g:    [][]int32{{0, 0, 1}, {1, 0, 1}},
			want: 0.50000000,
		},
		{
			name: "test 2",
			r:    2,
			c:    2,
			g:    [][]int32{{1, 1}, {1, 1}},
			want: 1.00000000,
		},
	}

	for _, test := range tests {
		have := getHitProbability(test.r, test.c, test.g)
		assert.Equalf(t, test.want, have, "%s: failed calculating probability",
			test.name)
	}
}
