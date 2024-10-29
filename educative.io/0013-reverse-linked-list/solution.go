package educative0013

import "github.com/bharath23/coding-go/internal"

func reverse(head *internal.ListNode) *internal.ListNode {
	var prev *internal.ListNode
	for cur := head; cur != nil; {
		tmp := cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}
