package leetcode0051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "test 1",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "test 2",
			n:    1,
			want: [][]string{{"Q"}},
		},
	}

	for _, test := range tests {
		have := solveNQueens(test.n)
		assert.ElementsMatchf(t, test.want, have, "%s: solveNQueens failed",
			test.name)
	}
}
