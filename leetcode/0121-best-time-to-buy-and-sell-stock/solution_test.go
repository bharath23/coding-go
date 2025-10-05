package leetcode0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "test 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "test 2",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, test := range tests {
		have := maxProfitv0(test.prices)
		assert.Equalf(t, test.want, have, "%s: maxPorfitv0 failed", test.name)
		have = maxProfitv1(test.prices)
		assert.Equalf(t, test.want, have, "%s: maxPorfitv1 failed", test.name)
	}
}
