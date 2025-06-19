package leetcode0863

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		root   [][]int
		target int
		k      int
		want   []int
	}{
		{
			name:   "test 1",
			root:   [][]int{{3}, {5}, {1}, {6}, {2}, {0}, {8}, nil, nil, {7}, {4}},
			target: 5,
			k:      2,
			want:   []int{7, 4, 1},
		},
		{
			name:   "test 2",
			root:   [][]int{{1}},
			target: 1,
			k:      3,
			want:   []int{},
		},
	}

	for _, test := range tests {
		root, target := sliceToBinaryTreeWithTarget(test.root, test.target)
		have := distanceK(root, target, test.k)
		assert.ElementsMatchf(t, test.want, have, "%s: distanceK failed",
			test.name)
	}
}

func sliceToBinaryTreeWithTarget(nodes [][]int, target int) (*TreeNode, *TreeNode) {
	if len(nodes) == 0 {
		return nil, nil
	}

	queue := []*TreeNode{}
	root := &TreeNode{Val: nodes[0][0]}
	queue = append(queue, root)
	var targetNode *TreeNode
	i := 1
	for i < len(nodes) {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}

		if nodes[i] != nil {
			node := &TreeNode{Val: nodes[i][0]}
			cur.Left = node
			queue = append(queue, node)
			if node.Val == target {
				targetNode = node
			}
		} else {
			queue = append(queue, nil)
		}

		i++
		if i < len(nodes) && nodes[i] != nil {
			node := &TreeNode{Val: nodes[i][0]}
			cur.Right = node
			queue = append(queue, node)
			if node.Val == target {
				targetNode = node
			}
		} else {
			queue = append(queue, nil)
		}

		i++
	}

	return root, targetNode
}
