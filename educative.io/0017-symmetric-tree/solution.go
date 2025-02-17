package edutcative0017

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

type queue[T constraints.Ordered] struct {
	elements []*internal.TreeNode[T]
	size     int
}

func (q *queue[T]) dequeue() (*internal.TreeNode[T], bool) {
	if q.size == 0 {
		return nil, false
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return e, true
}

func (q *queue[T]) enqueue(n *internal.TreeNode[T]) {
	q.elements = append(q.elements, n)
	q.size++
}

func (q *queue[T]) isEmpty() bool {
	return q.size == 0
}

func newQueue[T constraints.Ordered]() *queue[T] {
	q := &queue[T]{}
	q.elements = make([]*internal.TreeNode[T], 0)
	return q
}

func isSymmetric(root *internal.TreeNode[int]) bool {
	if root == nil {
		return true
	}

	q := newQueue[int]()
	q.enqueue(root.Left)
	q.enqueue(root.Right)
	for !q.isEmpty() {
		left, _ := q.dequeue()
		right, _ := q.dequeue()
		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		q.enqueue(left.Left)
		q.enqueue(right.Right)
		q.enqueue(left.Right)
		q.enqueue(right.Left)
	}

	return true
}
