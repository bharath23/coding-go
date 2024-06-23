package leetcode0019

import "github.com/bharath23/coding-go/internal"

/*
Simple two pass solution. First pass count the number of nodes in the list. In
the second pass find the node that needs to be deleted. Time complexity is O(n)
and space complexity is O(1).

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func removeNthFromEndV0(head *internal.ListNode, n int) *internal.ListNode {
	listLen := 0
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	var prev *internal.ListNode
	count := 0
	if head != nil {
		count++
	}

	cur = head
	for count < (listLen+1-n) && cur != nil {
		prev = cur
		cur = cur.Next
		count++
	}

	if prev != nil {
		prev.Next = cur.Next
	}

	if head == cur {
		head = cur.Next
	}

	return head
}

/*
Two pointer solution. Move the first pointer by n nodes. Then move first and
second pointer till we reach end of the list. We delete the node after the node
pointed to by the second pointer. Time complexity is still O(n) as we traverse
the list once. Space complexity is still O(1) as no additional storage is
required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func removeNthFromEndV1(head *internal.ListNode, n int) *internal.ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	var p2 *internal.ListNode
	for p1 != nil {
		p1 = p1.Next
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = head
		}
	}

	if p2 != nil {
		p2.Next = p2.Next.Next
	} else {
		head = head.Next
	}

	return head
}
