package meta0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		d    []int
		k    int
		want int
	}{
		{
			name: "test 1",
			n:    6,
			d:    []int{1, 2, 3, 3, 2, 1},
			k:    1,
			want: 5,
		},
		{
			name: "test 2",
			n:    6,
			d:    []int{1, 2, 3, 3, 2, 1},
			k:    2,
			want: 4,
		},
		{
			name: "test 1",
			n:    7,
			d:    []int{1, 2, 1, 2, 1, 2, 1},
			k:    2,
			want: 2,
		},
	}

	for _, test := range tests {
		have := getMaximumDishEatenCount(test.n, test.d, test.k)
		assert.Equalf(t, test.want, have, "%s: getMaximumEatenDishCount failed",
			test.name)
	}
}
