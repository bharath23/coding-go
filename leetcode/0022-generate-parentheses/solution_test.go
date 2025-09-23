package leetcode0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "test 1",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "test 2",
			n:    1,
			want: []string{"()"},
		},
	}

	for _, test := range tests {
		have := generateParenthesis(test.n)
		assert.ElementsMatchf(t, test.want, have, "%s: generateParenthesis failed",
			test.name)
	}
}
