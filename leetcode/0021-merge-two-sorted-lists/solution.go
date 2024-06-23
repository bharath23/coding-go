package leetcode0021

import "github.com/bharath23/coding-go/internal"

/*
Simple recursive solution. Time complexity is O(n) as we traverse both list
once. Space complexity is O(1) as no additional storage is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func mergeTwoListsV0(list1 *internal.ListNode, list2 *internal.ListNode) *internal.ListNode {
	var recursiveMergeLists func(*internal.ListNode, *internal.ListNode) *internal.ListNode
	recursiveMergeLists = func(l1, l2 *internal.ListNode) *internal.ListNode {
		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}

		if l1.Val <= l2.Val {
			l1.Next = recursiveMergeLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = recursiveMergeLists(l1, l2.Next)
			return l2
		}
	}

	return recursiveMergeLists(list1, list2)
}

/*
Simple list traversal solution. Compare the values at the head of the lists
and update lists accordingly. Time complexity O(n) as the lists are traversed
once. Space complexity is O(1) as no extra storage is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func mergeTwoListsV1(list1 *internal.ListNode, list2 *internal.ListNode) *internal.ListNode {
	var head, prev *internal.ListNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			if prev != nil {
				prev.Next = list1
			} else {
				head = list1
			}
			prev = list1
			list1 = list1.Next
		} else {
			if prev != nil {
				prev.Next = list2
			} else {
				head = list2
			}
			prev = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		if prev != nil {
			prev.Next = list1
		} else {
			head = list1
		}
	}

	if list2 != nil {
		if prev != nil {
			prev.Next = list2
		} else {
			head = list2
		}
	}

	return head
}
