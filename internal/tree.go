package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewListFromTree(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	q := NewQueue()
	q.Enqueue(root)
	for !q.IsEmpty() {
		node := q.Dequeue().(*TreeNode)
		if node == nil {
			continue
		}
		result = append(result, []int{node.Val})
		if node.Left != nil {
			q.Enqueue(node.Left)
		}

		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}

	return result
}

func NewTreeFromList(input [][]int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	q := NewQueue()
	root := &TreeNode{Val: input[0][0]}
	q.Enqueue(root)
	i := 1
	for i < len(input) {
		cur := q.Dequeue().(*TreeNode)
		if cur == nil {
			continue
		}

		if input[i] != nil {
			node := &TreeNode{Val: input[i][0]}
			cur.Left = node
			q.Enqueue(node)
		} else {
			q.Enqueue(nil)
		}

		i = i + 1
		if i < len(input) && input[i] != nil {
			node := &TreeNode{Val: input[i][0]}
			cur.Right = node
			q.Enqueue(node)
		} else {
			q.Enqueue(nil)
		}

		i = i + 1
	}

	return root
}
