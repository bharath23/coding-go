package meta0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		c    string
		want string
	}{
		{
			name: "test 1",
			n:    3,
			c:    "ABA",
			want: "BAB",
		},
		{
			name: "test 2",
			n:    5,
			c:    "BBBBB",
			want: "AAAAA",
		},
	}

	for _, test := range tests {
		have := getWrongAnswers(test.n, test.c)
		assert.Equalf(t, test.want, have, "%s: incorrect generation", test.name)
	}
}
