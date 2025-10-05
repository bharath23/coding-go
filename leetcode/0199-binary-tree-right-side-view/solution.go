package leetcode0199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * The solution performs a level-order traversal (BFS) of the binary tree.
 * 1. The outer loop iterates once per level of the tree.
 * 2. The inner loop processes all nodes at the current level.
 *
 * Since each node is visited exactly once, the total time complexity is O(n),
 * where n is the number of nodes in the tree.
 *
 * Additional space is used to store nodes in the queue (next_level and
 * cur_level), which in the worst case can hold up to O(n) nodes for the bottom
 * level. The result slice also stores up to O(h) elements, where h is the
 * height of the tree. Thus, the overall space complexity is O(n).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(n)
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	next_level := []*TreeNode{}
	next_level = append(next_level, root)
	for len(next_level) != 0 {
		cur_level := next_level
		next_level = []*TreeNode{}
		var cur *TreeNode
		for len(cur_level) != 0 {
			cur = cur_level[0]
			cur_level = cur_level[1:]
			if cur.Left != nil {
				next_level = append(next_level, cur.Left)
			}

			if cur.Right != nil {
				next_level = append(next_level, cur.Right)
			}
		}

		result = append(result, cur.Val)
	}

	return result
}
