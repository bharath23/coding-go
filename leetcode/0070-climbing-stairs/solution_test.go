package leetcode0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test 1",
			n:    2,
			want: 2,
		},
		{
			name: "test 2",
			n:    3,
			want: 3,
		},
		{
			name: "test 3",
			n:    45,
			want: 1836311903,
		},
	}

	for _, test := range tests {
		have := climbStairsV0(test.n)
		assert.Equalf(t, test.want, have, "%s: climbStair results incorrect", test.name)
	}

	for _, test := range tests {
		have := climbStairsV1(test.n)
		assert.Equalf(t, test.want, have, "%s: climbStair results incorrect", test.name)
	}

	for _, test := range tests {
		have := climbStairsV2(test.n)
		assert.Equalf(t, test.want, have, "%s: climbStair results incorrect", test.name)
	}
}
