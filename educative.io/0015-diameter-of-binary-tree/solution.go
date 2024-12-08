package educative0015

import (
	"github.com/bharath23/coding-go/internal"
)

func diameterOfBinaryTree(root *internal.TreeNode) int {
	maxDiameter := 0
	var findMaxDiameter func(*internal.TreeNode) int
	findMaxDiameter = func(node *internal.TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := findMaxDiameter(node.Left)
		rightHeight := findMaxDiameter(node.Right)
		maxDiameter = max(maxDiameter, leftHeight+rightHeight)
		return max(leftHeight, rightHeight) + 1
	}

	findMaxDiameter(root)
	return maxDiameter
}
