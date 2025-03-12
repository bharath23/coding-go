package educative0018

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

type element[T constraints.Ordered] struct {
	node  *internal.TreeNode[T]
	index int
}

type queue[T constraints.Ordered] struct {
	elements []*element[T]
	size     int
}

func (q *queue[T]) dequeue() (*element[T], bool) {
	if q.size == 0 {
		return nil, false
	}

	e := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return e, true
}

func (q *queue[T]) enqueue(n *internal.TreeNode[T], index int) {
	q.elements = append(q.elements, &element[T]{n, index})
	q.size++
}

func (q *queue[T]) isEmpty() bool {
	return q.size == 0
}

func newQueue[T constraints.Ordered]() *queue[T] {
	q := &queue[T]{}
	q.elements = make([]*element[T], 0)
	return q
}

func verticalOrder(root *internal.TreeNode[int]) [][]int {
	if root == nil {
		return nil
	}

	orderMap := make(map[int][]*internal.TreeNode[int])
	minIdx := 0
	maxIdx := 0
	q := newQueue[int]()
	q.enqueue(root, 0)
	for !q.isEmpty() {
		e, _ := q.dequeue()
		n := e.node
		idx := e.index
		if n.Left != nil {
			q.enqueue(n.Left, idx-1)
		}

		if n.Right != nil {
			q.enqueue(n.Right, idx+1)
		}
		minIdx = min(minIdx, idx)
		maxIdx = max(maxIdx, idx)
		o := orderMap[idx]
		if o == nil {
			o = make([]*internal.TreeNode[int], 0)
		}
		o = append(o, n)
		orderMap[idx] = o
	}

	result := make([][]int, 0)
	idx := 0
	for i := minIdx; i <= maxIdx; i++ {
		result = append(result, []int{})
		for _, n := range orderMap[i] {
			result[idx] = append(result[idx], n.Val)
		}
		idx++
	}

	return result
}
