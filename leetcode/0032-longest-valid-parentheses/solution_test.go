package leetcode0032

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "test 1",
			s:    "(()",
			want: 2,
		},
		{
			name: "test 2",
			s:    ")()())",
			want: 4,
		},
		{
			name: "test 3",
			s:    "",
			want: 0,
		},
		{
			name: "test 4",
			s:    "())((())",
			want: 4,
		},
	}

	for _, test := range tests {
		have := longestValidParenthesesv0(test.s)
		assert.Equalf(t, test.want, have, "%s: longestValidParenthesesv0 failed",
			test.name)
		have = longestValidParenthesesv1(test.s)
		assert.Equalf(t, test.want, have, "%s: longestValidParenthesesv1 failed",
			test.name)
		have = longestValidParenthesesv2(test.s)
		assert.Equalf(t, test.want, have, "%s: longestValidParenthesesv2 failed",
			test.name)
	}
}
