package leetcode0680

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
			s:    "aba",
			want: true,
		},
		{
			name: "test 2",
			s:    "abca",
			want: true,
		},
		{
			name: "test 3",
			s:    "abc",
			want: false,
		},
	}

	for _, test := range tests {
		have := validPalindrome(test.s)
		assert.Equalf(t, test.want, have, "%s: validPalindrome failed", test.name)
	}
}
