package leetcode2236

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

func checkTree[T constraints.Ordered](root *internal.TreeNode[T]) bool {
	return root.Val == (root.Left.Val + root.Right.Val)
}
