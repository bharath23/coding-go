package leetcode0050

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			name: "test 1",
			x:    2.00000,
			n:    10,
			want: 1024.00000,
		},
		{
			name: "test 2",
			x:    2.10000,
			n:    3,
			want: 9.26100,
		},
		{
			name: "test 3",
			x:    2.00000,
			n:    -2,
			want: 0.25,
		},
		{
			name: "test 4",
			x:    1.00000,
			n:    2147483647,
			want: 1.00000,
		},
	}

	for _, test := range tests {
		have := myPowv0(test.x, test.n)
		assert.InDeltaf(t, test.want, have, 0.000001, "%s: myPowv0 failed", test.name)
	}

	for _, test := range tests {
		have := myPowv1(test.x, test.n)
		assert.InDeltaf(t, test.want, have, 0.000001, "%s: myPowv0 failed", test.name)
	}
}
