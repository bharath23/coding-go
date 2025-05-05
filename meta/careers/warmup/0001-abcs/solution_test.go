package meta0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{
			name: "test 1",
			a:    1,
			b:    2,
			c:    3,
			want: 6,
		},
		{
			name: "test 2",
			a:    100,
			b:    100,
			c:    100,
			want: 300,
		},
		{
			name: "test 3",
			a:    85,
			b:    16,
			c:    93,
			want: 194,
		},
	}

	for _, test := range tests {
		have := getSum(test.a, test.b, test.c)
		assert.Equalf(t, test.want, have, "%s: test case failed", test.name)
	}
}
