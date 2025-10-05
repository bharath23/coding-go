package leetcode0114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * The solution performs a preorder traversal using an explicit stack.
 * 1. Each node is pushed to and popped from the stack exactly once.
 * 2. During traversal, the left and right pointers are updated to form a
 *    flattened linked list in-place.
 *
 * Since every node is visited once, the total time complexity is O(n),
 * where n is the number of nodes in the binary tree.
 *
 * The stack can hold up to O(h) nodes at a time, where h is the height
 * of the tree (O(n) in the worst case for a skewed tree).
 * Thus, the overall space complexity is O(n).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(n)
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

		cur.Left = nil
		if prev != nil {
			prev.Right = cur
		}
		prev = cur
	}
}
