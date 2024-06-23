package leetcode0061

import (
	"github.com/bharath23/coding-go/internal"
)

/*
Simple two pass solution. First pass create a cycle in the list and count the
number of nodes. Break the cycle at the new head by traversring
length - k%length nodes. Time complexity is O(n) and space complexity is O(1).

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func rotateRight(head *internal.ListNode, k int) *internal.ListNode {
	if head == nil {
		return nil
	}

	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	tail.Next = head
	cur := head
	for i := 0; i < length-(k%length)-1; i++ {
		cur = cur.Next
	}

	head = cur.Next
	cur.Next = nil

	return head
}
