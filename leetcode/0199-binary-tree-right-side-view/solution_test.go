package leetcode0199

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		root [][]int
		want []int
	}{
		{
			name: "test 1",
			root: [][]int{{1}, {2}, {3}, nil, {5}, nil, {4}},
			want: []int{1, 3, 4},
		},
		{
			name: "test 2",
			root: [][]int{{1}, {2}, {3}, {4}, nil, nil, nil, {5}},
			want: []int{1, 3, 4, 5},
		},
		{
			name: "test 3",
			root: [][]int{{1}, nil, {3}},
			want: []int{1, 3},
		},
		{
			name: "test 4",
			root: [][]int{},
			want: []int{},
		},
	}

	for _, test := range tests {
		root := sliceToTree(test.root)
		have := rightSideView(root)
		assert.Equalf(t, test.want, have, "%s: rightSideView failed", test.name)
	}
}

func sliceToTree(nodes [][]int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	queue := []*TreeNode{}
	root := &TreeNode{Val: nodes[0][0]}
	queue = append(queue, root)
	i := 1
	for i < len(nodes) {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}

		if nodes[i] != nil {
			n := &TreeNode{Val: nodes[i][0]}
			cur.Left = n
			queue = append(queue, n)
		} else {
			queue = append(queue, nil)
		}

		i++
		if i < len(nodes) && nodes[i] != nil {
			n := &TreeNode{Val: nodes[i][0]}
			cur.Right = n
			queue = append(queue, n)
		} else {
			queue = append(queue, nil)
		}

		i++
	}

	return root
}
