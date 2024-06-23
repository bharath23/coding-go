package leetcode0226

import (
	"github.com/bharath23/coding-go/internal"
)

func invertTreeV0(root *internal.TreeNode) *internal.TreeNode {
	if root != nil {
		root.Left, root.Right = invertTreeV0(root.Right), invertTreeV0(root.Left)
	}

	return root
}

func invertTreeV1(root *internal.TreeNode) *internal.TreeNode {
	if root == nil {
		return nil
	}

	var result *internal.TreeNode
	q := []*internal.TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if result == nil {
			result = node
		}

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return result
}
