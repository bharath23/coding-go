package educative0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test1",
			s:    "()",
			want: "()",
		},
		{
			name: "test2",
			s:    "(",
			want: "",
		},
		{
			name: "test3",
			s:    ")",
			want: "",
		},
		{
			name: "test4",
			s:    ")(",
			want: "",
		},
		{
			name: "test5",
			s:    "ab)ca(so)(sc(s)(",
			want: "abca(so)sc(s)",
		},
		{
			name: "test6",
			s:    ")((xyz)())",
			want: "((xyz)())",
		},
	}

	for _, test := range tests {
		have := minRemoveParentheses(test.s)
		assert.Equalf(t, test.want, have, "%s: minRemoveParentheses incorrect result", test.name)
	}
}
