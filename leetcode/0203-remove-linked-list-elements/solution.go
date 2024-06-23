package leetcode0203

import "github.com/bharath23/coding-go/internal"

/*
Simple one pass solution. Check if the value of the node matches the value
which needs to be removed, if so delete the node. Time complexity is O(n) as
we traverse the entire list once. Space complexity is O(1) as we do not need
any additional storage.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func removeElements(head *internal.ListNode, val int) *internal.ListNode {
	var prev *internal.ListNode
	cur := head
	for cur != nil {
		if cur.Val == val {
			if prev != nil {
				prev.Next = cur.Next
			}

			if head == cur {
				head = cur.Next
			}
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return head
}
