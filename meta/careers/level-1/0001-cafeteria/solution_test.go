package meta0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		m    int32
		s    []int
		want int
	}{
		{
			name: "test 1",
			n:    10,
			k:    1,
			m:    2,
			s:    []int{2, 6},
			want: 3,
		},
		{
			name: "test 2",
			n:    15,
			k:    2,
			m:    3,
			s:    []int{11, 6, 14},
			want: 1,
		},
	}

	for _, test := range tests {
		have := getMaxAdditionalDinersCount(test.n, test.k, test.m, test.s)
		assert.Equalf(t, test.want, have,
			"%s: getMaxAdditionalDinersCount failed", test.name)
	}
}
