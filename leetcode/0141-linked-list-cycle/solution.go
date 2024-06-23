package leetcode0141

import "github.com/bharath23/coding-go/internal"

/*
Simple one pass solution. We use hash map to store all visited nodes. We check
the hash map to see if we have already visited a node or not. Time complexity
is O(n) as we traverse the list once. Space complexity is O(n) as we need to
maintain a hash map.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func hasCycleV0(head *internal.ListNode) bool {
	visited := map[*internal.ListNode]bool{}
	cur := head
	for cur != nil {
		isVisited := visited[cur]
		if isVisited {
			return true
		}

		visited[cur] = true
		cur = cur.Next
	}

	return false
}

/*
Two pointer solution. We use two pointers, a slow pointer and a fast pointer.
The fast pointer moves twice as fast as the slow pointer. If there is a cycle
the slow pointer will catch up to the fast pointer. If there is no cycle the
fast pointer will come to the end of the list. Time complexity is still O(n)
as we might have to visit all nodes once. Space complexity is O(1) as no
additional storage is needed.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func hasCycleV1(head *internal.ListNode) bool {
	var p1, p2 *internal.ListNode
	if head == nil {
		return false
	}

	p1 = head
	p2 = head
	for (p2 != nil) && (p2.Next != nil) {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}

	return false
}
