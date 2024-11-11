package educative0014

import (
	"github.com/bharath23/coding-go/internal"
)

func reverseKGroups(head *internal.ListNode, k int) *internal.ListNode {
	var prev *internal.ListNode
	for cur := head; cur != nil; {
		kStart := cur
		kEnd := cur
		count := 0
		for cur != nil && count < k {
			count++
			kEnd = cur
			cur = cur.Next
		}

		if count < k {
			break
		}

		kEnd.Next = nil
		var kPrev *internal.ListNode
		for kCur := kStart; kCur != nil; {
			tmp := kCur
			kCur = kCur.Next
			tmp.Next = kPrev
			kPrev = tmp
		}

		if prev != nil {
			prev.Next = kEnd
		} else {
			head = kEnd
		}

		kStart.Next = cur
		prev = kStart
	}

	return head
}
