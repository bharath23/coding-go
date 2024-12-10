package leetcode0876

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

func middleNode[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	cur := head
	mid := head

	for cur != nil && cur.Next != nil {
		mid = mid.Next
		cur = cur.Next.Next
	}

	return mid
}
