package meta0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		c    []int
		want int
	}{
		{
			name: "test 1",
			m:    3,
			n:    3,
			c:    []int{1, 2, 3},
			want: 2,
		},
		{
			name: "test 2",
			m:    4,
			n:    10,
			c:    []int{9, 4, 4, 8},
			want: 11,
		},
	}

	for _, test := range tests {
		have := getMinCodeEntryTime(test.n, test.m, test.c)
		assert.Equalf(t, test.want, have, "%s: getMinCodeEntryTime failed\n",
			test.name)
	}
}
