package leetcode0408

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		word string
		abbr string
		want bool
	}{
		{
			name: "test 1",
			word: "internationalization",
			abbr: "i12iz4n",
			want: true,
		},
		{
			name: "test 2",
			word: "apple",
			abbr: "a2e",
			want: false,
		},
	}

	for _, test := range tests {
		have := validWordAbbreviation(test.word, test.abbr)
		assert.Equalf(t, test.want, have, "%s: validWordAbbreviation failed", test.name)
	}
}
