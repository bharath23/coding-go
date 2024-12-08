package educative0015

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name:  "test 1",
			input: [][]int{{1}, {2}, {3}, {4}, {5}, {6}},
			want:  4,
		},
		{
			name:  "test 2",
			input: [][]int{{12}, {7}, {1}, {9}, {10}, {15}},
			want:  4,
		},
		{
			name:  "test 3",
			input: [][]int{{3}, {9}, {20}, {15}, {7}},
			want:  3,
		},
		{
			name:  "test 4",
			input: [][]int{{1}, {2}},
			want:  1,
		},
		{
			name:  "test 5",
			input: [][]int{{9}, {7}, nil, nil, {1}, {8}, {10}, nil, {12}},
			want:  4,
		},
	}

	for _, test := range tests {
		root := internal.NewTreeFromList(test.input)
		have := diameterOfBinaryTree(root)
		assert.Equalf(t, test.want, have, "%s: mismatch in diameter",
			test.name)
	}
}
