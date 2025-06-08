package leetcode0023

type ListNode struct {
	val  int
	next *ListNode
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	var head, cur, n *ListNode
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			n = l1
			l1 = l1.next
		} else {
			n = l2
			l2 = l2.next
		}

		n.next = nil
		if head == nil {
			head = n
		}

		if cur != nil {
			cur.next = n
		}

		cur = n
	}

	if cur != nil {
		if l1 != nil {
			cur.next = l1
		}

		if l2 != nil {
			cur.next = l2
		}
	} else {
		if l1 != nil {
			head = l1
		}

		if l2 != nil {
			head = l2
		}
	}

	return head
}

/*
We can merge two lists in O(n) time, where n is the total number of nodes in
two lists. Hence total time complexity is O(Nlogk), N is the total number of
nodes, k is the number of lists. The space complexity is O(1) as we do not
allocate any extra space.

Time complexity: O(Nlogk)
Space complexity: O(1)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	m := len(lists) / 2
	return mergeLists(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}
