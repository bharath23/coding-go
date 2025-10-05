package leetcode0199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
