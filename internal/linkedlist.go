package internal

import "golang.org/x/exp/constraints"

type ListNode[T constraints.Integer] struct {
	Val  T
	Next *ListNode[T]
}

func NewCircularListFromSlice[T constraints.Integer](input []T) *ListNode[T] {
	var head, tail *ListNode[T]
	for _, v := range input {
		n := &ListNode[T]{Val: v}
		if head == nil {
			head = n
		}

		n.Next = head
		if tail == nil {
			tail = n
		}

		tail.Next = n
		tail = n
	}

	return head
}

func NewListFromSlice[T constraints.Integer](input []T) *ListNode[T] {
	var head, tail *ListNode[T]
	for _, v := range input {
		n := &ListNode[T]{Val: v}
		if head == nil {
			head = n
		}

		if tail != nil {
			tail.Next = n
		}

		tail = n
	}

	return head
}

func NewListFromSliceWithCycle[T constraints.Integer](input []T, pos int) (
	*ListNode[T],
	*ListNode[T],
) {
	var cycle, head, tail *ListNode[T]
	for i, v := range input {
		n := &ListNode[T]{Val: v}
		if head == nil {
			head = n
		}

		if tail != nil {
			tail.Next = n
		}

		tail = n
		if i == pos {
			cycle = n
		}
	}

	if tail != nil {
		tail.Next = cycle
	}

	return head, cycle
}

func NewListsWithIntersection[T constraints.Integer](
	val T,
	skipA, skipB int,
	inputA, inputB []T,
) (
	*ListNode[T],
	*ListNode[T],
	*ListNode[T]) {
	var headA, tailA, headB, tailB, intersect *ListNode[T]
	for i, v := range inputA {
		n := &ListNode[T]{Val: v}
		if headA == nil {
			headA = n
		}

		if tailA != nil {
			tailA.Next = n
		}

		tailA = n
		if i == skipA {
			intersect = n
		}
	}

	for i, v := range inputB {
		if (i == skipB) && (v == val) {
			if headB == nil {
				headB = intersect
			}

			if tailB != nil {
				tailB.Next = intersect
			}

			tailB = tailA
		}

		n := &ListNode[T]{Val: v}
		if headB == nil {
			headB = n
		}

		if tailB != nil {
			tailB.Next = n
		}

		tailB = n
	}
	return headA, headB, intersect
}

func ListToSlice[T constraints.Integer](head *ListNode[T]) []T {
	s := []T{}
	for cur := head; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}

	return s
}

func CircularListToSlice[T constraints.Integer](head *ListNode[T]) []T {
	s := []T{}
	cur := head
	if cur != nil {
		s = append(s, cur.Val)
		cur = cur.Next
	}

	for cur != nil && cur != head {
		s = append(s, cur.Val)
		cur = cur.Next
	}

	return s
}
