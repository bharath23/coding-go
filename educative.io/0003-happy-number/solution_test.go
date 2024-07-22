package educative0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "test1",
			num:  2147483646,
			want: false,
		},
		{
			name: "test2",
			num:  1,
			want: true,
		},
		{
			name: "test3",
			num:  19,
			want: true,
		},
		{
			name: "test4",
			num:  8,
			want: false,
		},
		{
			name: "test5",
			num:  7,
			want: true,
		},
	}

	for _, test := range tests {
		have := isHappy(test.num)
		assert.Equalf(t, test.want, have, "%s: isHappy incorrect result", test.name)
	}
}
