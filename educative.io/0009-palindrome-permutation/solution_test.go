package educative0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		st   string
		want bool
	}{
		{
			name: "test1",
			st:   "abab",
			want: true,
		},
		{
			name: "test2",
			st:   "peas",
			want: false,
		},
		{
			name: "test3",
			st:   "racecar",
			want: true,
		},
		{
			name: "test4",
			st:   "code",
			want: false,
		},
		{
			name: "test5",
			st:   "baefeab",
			want: true,
		},
	}

	for _, test := range tests {
		have := permutatePalindrome(test.st)
		assert.Equalf(t, test.want, have, "%s: permutatePalindrome result incorrect", test.name)
	}
}
