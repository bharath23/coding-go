package leetcode0328

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Simple one pass solution. As we traverse the list delete all even nodes and
add it to a separate list. Append the even list to the tail of the odd list.
Time complexity is O(n) as we need to traverse te list once. Space complexity
is O(1) as no additional space is required.

Complexity:
Time complexity; O(n) each node needs to be traversed just once
Space complexity: O(1) no addiional storage is required
*/
func oddEvenListV0[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	var prev, headEven, tailEven *internal.ListNode[T]
	i := 1
	cur := head
	for cur != nil {
		if i%2 == 0 {
			if prev != nil {
				prev.Next = cur.Next
			}

			if head == cur {
				head = cur.Next
			}
			tmp := cur
			cur = cur.Next
			tmp.Next = nil
			if headEven == nil {
				headEven = tmp
			}

			if tailEven != nil {
				tailEven.Next = tmp
			}

			tailEven = tmp
		} else {
			prev = cur
			cur = cur.Next
		}

		i++
	}

	if prev != nil {
		prev.Next = headEven
	}

	if head == nil {
		head = headEven
	}

	return head
}

/*
Simple two pointer solution. We have two pointers even and odd to traverse the
list. Even pointer acts as the tail pointer for the even list and odd pointer
for the odd list. Time complexity is O(n) since we traverse the list once.
Space complexity is O(1) as we do not need any additional storage.

Complexity:
Time complexity: O(n), list traversed once.
Space complexity: O(1), no additional storage required
*/
func oddEvenListV1[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	headEven := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = headEven
	return head
}
