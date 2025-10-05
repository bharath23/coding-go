package leetcode0114

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		nodes [][]int
		want  [][]int
	}{
		{
			name:  "test 1",
			nodes: [][]int{{1}, {2}, {5}, {3}, {4}, nil, {6}},
			want:  [][]int{{1}, nil, {2}, nil, {3}, nil, {4}, nil, {5}, nil, {6}},
		},
		{
			name:  "test 2",
			nodes: [][]int{},
			want:  [][]int{},
		},
		{
			name:  "test 1",
			nodes: [][]int{{0}},
			want:  [][]int{{0}},
		},
	}

	for _, test := range tests {
		root := sliceToTree(test.nodes)
		flatten(root)
		have := treeToSlice(root)
		assert.Equalf(t, test.want, have, "%s: flatten failed", test.name)
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

func treeToSlice(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	slice := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			slice = append(slice, nil)
			continue
		}

		slice = append(slice, []int{node.Val})
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	i := len(slice) - 1
	for i >= 0 && slice[i] == nil {
		i--
	}
	return slice[:i+1]
}
