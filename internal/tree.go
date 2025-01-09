package internal

import "golang.org/x/exp/constraints"

type TreeNode[T constraints.Ordered] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type queue[T constraints.Ordered] struct {
	elements []*TreeNode[T]
	size     int
}

func (q *queue[T]) dequeue() (*TreeNode[T], bool) {
	if q.size == 0 {
		return nil, false
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return e, true
}

func (q *queue[T]) enqueue(n *TreeNode[T]) {
	q.elements = append(q.elements, n)
	q.size++
}

func (q *queue[T]) isEmpty() bool {
	return q.size == 0
}

func newQueue[T constraints.Ordered]() *queue[T] {
	q := &queue[T]{}
	q.elements = make([]*TreeNode[T], 0)
	return q
}

func NewListFromTree[T constraints.Ordered](root *TreeNode[T]) [][]T {
	if root == nil {
		return [][]T{}
	}

	result := [][]T{}
	q := newQueue[T]()
	q.enqueue(root)
	for !q.isEmpty() {
		node, _ := q.dequeue()
		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, []T{node.Val})
		q.enqueue(node.Left)
		q.enqueue(node.Right)
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func NewTreeFromList[T constraints.Ordered](input [][]T) *TreeNode[T] {
	if len(input) == 0 {
		return nil
	}

	q := newQueue[T]()
	root := &TreeNode[T]{Val: input[0][0]}
	q.enqueue(root)
	i := 1
	for i < len(input) {
		cur, _ := q.dequeue()
		if cur == nil {
			continue
		}

		if input[i] != nil {
			node := &TreeNode[T]{Val: input[i][0]}
			cur.Left = node
			q.enqueue(node)
		} else {
			q.enqueue(nil)
		}

		i = i + 1
		if i < len(input) && input[i] != nil {
			node := &TreeNode[T]{Val: input[i][0]}
			cur.Right = node
			q.enqueue(node)
		} else {
			q.enqueue(nil)
		}

		i = i + 1
	}

	return root
}
