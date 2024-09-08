package educative0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name        string
		numerator   int
		denominator int
		want        string
	}{
		{
			name:        "test1",
			numerator:   2,
			denominator: 4,
			want:        "0.5",
		},
		{
			name:        "test2",
			numerator:   4,
			denominator: 2,
			want:        "2",
		},
		{
			name:        "test3",
			numerator:   5,
			denominator: 333,
			want:        "0.(015)",
		},
		{
			name:        "test4",
			numerator:   2,
			denominator: 3,
			want:        "0.(6)",
		},
		{
			name:        "test5",
			numerator:   47,
			denominator: 18,
			want:        "2.6(1)",
		},
	}

	for _, test := range tests {
		have := fractionToDecimal(test.numerator, test.denominator)
		assert.Equalf(t, test.want, have, "%s: fractionToDecimal incorrect retval", test.name)
	}
}
