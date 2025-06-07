package leetcode0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "test 1",
			s:    "()",
			want: true,
		},
		{
			name: "test 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "test 3",
			s:    "(]",
			want: false,
		},
		{
			name: "test 4",
			s:    "([])",
			want: true,
		},
		{
			name: "test 5",
			s:    "(",
			want: false,
		},
	}

	for _, test := range tests {
		have := isValid(test.s)
		assert.Equalf(t, test.want, have, "%s: isValid failed", test.name)
	}
}
