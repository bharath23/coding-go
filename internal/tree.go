package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	elements []*TreeNode
	size     int
}

func (q *queue) dequeue() (*TreeNode, bool) {
	if q.size == 0 {
		return nil, false
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return e, true
}

func (q *queue) enqueue(n *TreeNode) {
	q.elements = append(q.elements, n)
	q.size++
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func newQueue() *queue {
	q := &queue{}
	q.elements = make([]*TreeNode, 0)
	return q
}

func NewListFromTree(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	q := newQueue()
	q.enqueue(root)
	for !q.isEmpty() {
		node, _ := q.dequeue()
		if node == nil {
			continue
		}

		result = append(result, []int{node.Val})
		if node.Left != nil {
			q.enqueue(node.Left)
		}

		if node.Right != nil {
			q.enqueue(node.Right)
		}
	}

	return result
}

func NewTreeFromList(input [][]int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	q := newQueue()
	root := &TreeNode{Val: input[0][0]}
	q.enqueue(root)
	i := 1
	for i < len(input) {
		cur, _ := q.dequeue()
		if cur == nil {
			continue
		}

		if input[i] != nil {
			node := &TreeNode{Val: input[i][0]}
			cur.Left = node
			q.enqueue(node)
		} else {
			q.enqueue(nil)
		}

		i = i + 1
		if i < len(input) && input[i] != nil {
			node := &TreeNode{Val: input[i][0]}
			cur.Right = node
			q.enqueue(node)
		} else {
			q.enqueue(nil)
		}

		i = i + 1
	}

	return root
}
