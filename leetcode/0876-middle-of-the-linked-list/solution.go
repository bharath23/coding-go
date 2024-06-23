package leetcode0876

import (
	"github.com/bharath23/coding-go/internal"
)

func middleNode(head *internal.ListNode) *internal.ListNode {
	cur := head
	mid := head

	for cur != nil && cur.Next != nil {
		mid = mid.Next
		cur = cur.Next.Next
	}

	return mid
}
