package leetcode0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "test 1",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "test 2",
			m:    3,
			n:    2,
			want: 3,
		},
	}

	for _, test := range tests {
		have := uniquePaths(test.m, test.n)
		assert.Equalf(t, test.want, have, "%s: uniquePaths failed", test.name)
	}
}
