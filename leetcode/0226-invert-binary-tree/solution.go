package leetcode0226

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

func invertTreeV0[T constraints.Ordered](root *internal.TreeNode[T]) *internal.TreeNode[T] {
	if root != nil {
		root.Left, root.Right = invertTreeV0(root.Right), invertTreeV0(root.Left)
	}

	return root
}

func invertTreeV1[T constraints.Ordered](root *internal.TreeNode[T]) *internal.TreeNode[T] {
	if root == nil {
		return nil
	}

	var result *internal.TreeNode[T]
	q := []*internal.TreeNode[T]{}
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
