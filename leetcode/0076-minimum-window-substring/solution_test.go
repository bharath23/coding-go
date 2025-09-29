package leetcode0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "test 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "test 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "test 1",
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for _, test := range tests {
		have := minWindowv0(test.s, test.t)
		assert.Equalf(t, test.want, have, "%s, minWindowv0 failed", test.name)
		have = minWindowv1(test.s, test.t)
		assert.Equalf(t, test.want, have, "%s, minWindowv1 failed", test.name)
	}
}
