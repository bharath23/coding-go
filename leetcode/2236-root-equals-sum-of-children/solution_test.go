package leetcode2236

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		nodes [][]int
		want  bool
	}{
		{
			name:  "test 1",
			nodes: [][]int{{10}, {4}, {6}},
			want:  true,
		},
		{
			name:  "test 2",
			nodes: [][]int{{5}, {3}, {1}},
			want:  false,
		},
	}

	for _, test := range tests {
		root := internal.NewTreeFromList(test.nodes)
		have := checkTree(root)
		assert.Equalf(t, test.want, have, "%s: checkTree result does not match", test.name)
	}
}
