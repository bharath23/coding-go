package leetcode0025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	cur := head
	for k > 0 {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
		k--
	}

	return newHead
}

/*
 * We process each node twice, once while counting to see if we have k nodes
 * and once when we reverse the sublist. The time complexity is O(n).
 * Space complexity is O(1) with and additional space complexity of O(n/k) for
 * the recursion stack.
 *
 * Time complexity: O(n)
 * Space complexity: O(1)
 *
 */
func reverseKGroupv0(head *ListNode, k int) *ListNode {
	count := 0
	cur := head
	for (count < k) && (cur != nil) {
		cur = cur.Next
		count++
	}

	if count == k {
		revHead := reverseLinkedList(head, k)
		head.Next = reverseKGroupv0(cur, k)
		return revHead
	}

	return head
}

/*
 * We process each node twice, once while counting to see if we have k nodes
 * and once when we reverse the sublist. The time complexity is O(n).
 * Space complexity is O(1) with and no additional space complexity.
 *
 * Time complexity: O(n)
 * Space complexity: O(1)
 *
 */
func reverseKGroupv1(head *ListNode, k int) *ListNode {
	newHead := head
	cur := head
	var tail *ListNode
	for cur != nil {
		count := 0
		for (count < k) && cur != nil {
			cur = cur.Next
			count++
		}
		if count == k {
			revHead := reverseLinkedList(head, k)
			if newHead == head {
				newHead = revHead
			}
			if tail != nil {
				tail.Next = revHead
			}

			tail = head
			head = cur
		}
	}

	if tail != nil {
		tail.Next = head
	}

	return newHead
}
