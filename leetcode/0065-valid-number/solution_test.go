package leetcode0065

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
			s:    "0",
			want: true,
		},
		{
			name: "test 2",
			s:    "e",
			want: false,
		},
		{
			name: "test 3",
			s:    ".",
			want: false,
		},
		{
			name: "test 4",
			s:    "e9",
			want: false,
		},
		{
			name: "test 5",
			s:    "-123.456E+789",
			want: true,
		},
	}

	for _, test := range tests {
		have := isNumberv0(test.s)
		assert.Equalf(t, test.want, have, "%s: isNumberv0 failed", test.name)
	}

	for _, test := range tests {
		have := isNumberv1(test.s)
		assert.Equalf(t, test.want, have, "%s: isNumberv1 failed", test.name)
	}
}
