package leetcode0002

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Iterate through the lists adding the values and carry-over as request. Once
either of the list ends process the remaning nodes

Complexity:
Time complexity: O(n), we traverse each element of the list once
Space complexity: O(n), we need extra space to store the new list
*/

func addTwoNumbers[T constraints.Integer](
	l1 *internal.ListNode[T],
	l2 *internal.ListNode[T],
) *internal.ListNode[T] {
	var l3, cur *internal.ListNode[T]
	var carry T
	carry = 0
	for (l1 != nil) || (l2 != nil) {
		v := carry
		if l1 != nil {
			v += l1.Val
		}

		if l2 != nil {
			v += l2.Val
		}

		carry = v / 10
		n := &internal.ListNode[T]{Val: v % 10}
		if l3 == nil {
			l3 = n
		}

		if cur != nil {
			cur.Next = n
		}

		cur = n
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		n := &internal.ListNode[T]{Val: carry}
		if cur != nil {
			cur.Next = n
		}

		if l3 == nil {
			l3 = n
		}
	}

	return l3
}
