package internal

import "golang.org/x/exp/constraints"

type Node[T constraints.Signed] struct {
	Val   T
	Child *Node[T]
	Next  *Node[T]
	Prev  *Node[T]
}

func NewMultiLevelDoublyLinkedList[T constraints.Signed](
	input []T,
) *Node[T] {
	newDoublyLinkedList := func(input []T, start, end int) (*Node[T], int) {
		var head, tail *Node[T]
		i := start
		for i < end {
			if input[i] == -1 {
				i++
				break
			}
			n := &Node[T]{Val: input[i]}
			if head == nil {
				head = n
			}

			n.Prev = tail
			if tail != nil {
				tail.Next = n
			}

			tail = n
			i++
		}

		return head, i
	}

	var head, cur *Node[T]
	end := len(input)
	i := 0
	for i < end {
		if input[i] == -1 {
			cur = cur.Next
			i++
		} else {
			var node *Node[T]
			node, i = newDoublyLinkedList(input, i, end)
			if head == nil {
				head = node
			}
			if cur != nil {
				cur.Child = node
			}

			cur = node
		}
	}

	return head
}

func MultiLevelDoublyLinkedListToSlice[T constraints.Signed](
	head *Node[T],
) []T {
	s := []T{}
	for cur := head; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}

	return s
}
