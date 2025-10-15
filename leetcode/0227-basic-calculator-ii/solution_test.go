package leetcode0227

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
			s:    "3+2*2",
			want: 7,
		},
		{
			name: "test 2",
			s:    " 3/2 ",
			want: 1,
		},
		{
			name: "test 3",
			s:    " 3+5 / 2 ",
			want: 5,
		},
	}

	for _, test := range tests {
		have := calculatev0(test.s)
		assert.Equalf(t, test.want, have, "%s: calculatev0 failed", test.name)
		have = calculatev1(test.s)
		assert.Equalf(t, test.want, have, "%s: calculatev1 failed", test.name)
	}
}
