package leetcode0042

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "test 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "test 2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, test := range tests {
		have := trapv0(test.height)
		assert.Equalf(t, test.want, have, "%s: trapv0 calculation failed", test.name)
	}

	for _, test := range tests {
		have := trapv1(test.height)
		assert.Equalf(t, test.want, have, "%s: trapv1 calculation failed", test.name)
	}
}
