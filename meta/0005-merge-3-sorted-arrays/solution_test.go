package meta0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		c    []int
		want []int
	}{
		{
			name: "test 1",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{2, 3, 4},
			c:    []int{4, 5, 6, 7},
			want: []int{1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 6, 7},
		},
		{
			name: "test 2",
			a:    []int{1, 2, 3, 5},
			b:    []int{6, 7, 8, 9},
			c:    []int{10, 11, 12},
			want: []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}

	for _, test := range tests {
		have := merge(test.a, test.b, test.c)
		assert.Equalf(t, test.want, have, "%s: merge failed", test.name)
	}
}
