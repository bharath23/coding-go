package educative0026

type LinkedListNode struct {
	val  int
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
}

func mergeKLists(lists []*LinkedList) *LinkedListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}

	if k == 1 {
		return lists[0].head
	}

	m := k / 2

	return merge(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}

func merge(list1, list2 *LinkedListNode) *LinkedListNode {
	var head, cur *LinkedListNode
	for list1 != nil && list2 != nil {
		if list1.val < list2.val {
			if head == nil {
				head = list1
			}

			if cur != nil {
				cur.next = list1
				cur = cur.next
			} else {
				cur = list1
			}

			list1 = list1.next
		} else {
			if head == nil {
				head = list2
			}

			if cur != nil {
				cur.next = list2
				cur = cur.next
			} else {
				cur = list2
			}

			list2 = list2.next
		}
	}

	if list1 != nil {
		cur.next = list1
	}

	if list2 != nil {
		cur.next = list2
	}

	return head
}
