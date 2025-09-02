package leetcode0029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		divident int32
		divisor  int32
		want     int32
	}{
		{
			name:     "test 1",
			divident: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "test 2",
			divident: 7,
			divisor:  -3,
			want:     -2,
		},
		{
			name:     "test 3",
			divident: 2147483647,
			divisor:  -1,
			want:     -2147483647,
		},
	}

	for _, test := range tests {
		have := dividev0(test.divident, test.divisor)
		assert.Equalf(t, test.want, have, "%s: divide failed", test.name)
	}

	for _, test := range tests {
		have := dividev1(test.divident, test.divisor)
		assert.Equalf(t, test.want, have, "%s: divide failed", test.name)
	}

	for _, test := range tests {
		have := dividev2(test.divident, test.divisor)
		assert.Equalf(t, test.want, have, "%s: divide failed", test.name)
	}
}
