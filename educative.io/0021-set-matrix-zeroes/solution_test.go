package educative0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{
			name: "test 1",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 0, 9}},
			want: [][]int{{1, 0, 3}, {4, 0, 6}, {0, 0, 0}},
		},
		{
			name: "test 2",
			mat:  [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {8, 9, 4, 6}},
			want: [][]int{{1, 2, 3, 4}, {4, 5, 6, 7}, {8, 9, 4, 6}},
		},
		{
			name: "test 3",
			mat: [][]int{{1, 1, 0, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}},
			want: [][]int{{0, 0, 0, 0, 0}, {1, 0, 0, 1, 1}, {1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0}, {1, 0, 0, 1, 1}},
		},
		{
			name: "test 4",
			mat:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name: "test 5",
			mat:  [][]int{{3, 5, 2, 0}, {1, 0, 4, 6}, {7, 3, 2, 4}},
			want: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {7, 0, 2, 0}},
		},
	}

	test := tests[0]
	have := setMatrixZeros(test.mat)
	assert.Equalf(t, test.want, have, "%s: incorrect setMatrixZeroes "+
		"return value", test.name)
}
