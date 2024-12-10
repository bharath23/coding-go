package educative0013

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

func reverse[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	var prev *internal.ListNode[T]
	for cur := head; cur != nil; {
		tmp := cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}
