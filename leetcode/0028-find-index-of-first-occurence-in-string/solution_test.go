package leetcode0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "test 1",
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			name:     "test 2",
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			name:     "test 3",
			haystack: "a",
			needle:   "a",
			want:     0,
		},
	}

	for _, test := range tests {
		have := strStrv0(test.haystack, test.needle)
		assert.Equalf(t, test.want, have, "%s: strStrv0 failed", test.name)
	}

	for _, test := range tests {
		have := strStrv1(test.haystack, test.needle)
		assert.Equalf(t, test.want, have, "%s: strStrv1 failed", test.name)
	}

	for _, test := range tests {
		have := strStrKMP(test.haystack, test.needle)
		assert.Equalf(t, test.want, have, "%s: strStrKMP failed", test.name)
	}
}
