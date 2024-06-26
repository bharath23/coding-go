package leetcode1137

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
			n:    4,
			want: 4,
		},
		{
			name: "test 2",
			n:    25,
			want: 1389537,
		},
		{
			name: "test 3",
			n:    35,
			want: 615693474,
		},
		{
			name: "test 4",
			n:    0,
			want: 0,
		},
	}

	for _, test := range tests {
		have := tribonacciV0(test.n)
		assert.Equalf(t, test.want, have, "%s: tribonacciV0 result mismatch", test.name)
	}

	for _, test := range tests {
		have := tribonacciV1(test.n)
		assert.Equalf(t, test.want, have, "%s: tribonacciV1 result mismatch", test.name)
	}

	for _, test := range tests {
		have := tribonacciV2(test.n)
		assert.Equalf(t, test.want, have, "%s: tribonacciV2 result mismatch", test.name)
	}
}
