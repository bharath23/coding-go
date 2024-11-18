package leetcode0226

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	nodes [][]int
	want  [][]int
}{
	{
		name:  "test 1",
		nodes: [][]int{{4}, {2}, {7}, {1}, {3}, {6}, {9}},
		want:  [][]int{{4}, {7}, {2}, {9}, {6}, {3}, {1}},
	},
	{
		name:  "test 2",
		nodes: [][]int{{2}, {1}, {3}},
		want:  [][]int{{2}, {3}, {1}},
	},
	{
		name:  "test 3",
		nodes: [][]int{},
		want:  [][]int{},
	},
}

func TestSolutionV0(t *testing.T) {

	for _, test := range tests {
		root := internal.NewTreeFromList(test.nodes)
		have := internal.NewListFromTree(invertTreeV0(root))
		assert.Equalf(t, test.want, have, "%s: invertTree result does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {

	for _, test := range tests {
		root := internal.NewTreeFromList(test.nodes)
		have := internal.NewListFromTree(invertTreeV1(root))
		assert.Equalf(t, test.want, have, "%s: invertTree result does not match", test.name)
	}
}
