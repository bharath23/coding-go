package leetcode0708

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Two pointer solution. We use two pointer solution to locate where in the list
we want insert the new value. Time complexity is O(n) as we could visit every
node in the list. Space complexity is O(1) as no additionla storae is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func insert[T constraints.Integer](
	aNode *internal.ListNode[T],
	x T,
) *internal.ListNode[T] {
	n := &internal.ListNode[T]{Val: x}
	if aNode == nil {
		n.Next = n
		return n
	}

	prev := aNode
	cur := aNode.Next
	for cur != aNode {
		if (prev.Val <= x) && (x <= cur.Val) {
			break
		} else if prev.Val > cur.Val {
			if (prev.Val <= x) || (x <= cur.Val) {
				break
			}
		}

		prev = cur
		cur = cur.Next
	}

	if cur.Val > x {
		n.Next = cur
		prev.Next = n
	} else if prev.Val <= x {
		n.Next = cur.Next
		cur.Next = n
	}

	return aNode
}
