package leetcode0141

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Simple one pass solution. We use hash map to store all visited nodes. We check
the hash map to see if we have already visited a node or not. Time complexity
is O(n) as we traverse the list once. Space complexity is O(n) as we need to
maintain a hash map.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func detectCycleV0[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	visited := map[*internal.ListNode[T]]bool{}
	cur := head
	for cur != nil {
		isVisited := visited[cur]
		if isVisited {
			return cur
		}

		visited[cur] = true
		cur = cur.Next
	}

	return nil
}

/*
Two pointer solution. We use two pointers, a slow pointer and a fast pointer.
The fast pointer moves twice as fast as the slow pointer. If there is a cycle
the slow pointer will catch up to the fast pointer. If there is no cycle the
fast pointer will come to the end of the list. Once we find the interseection
we reset the slow pointer to head. Move both the slow pointer and faster one
node at time till they are equal. Time complexity is still O(n) as we might
have to visit all nodes once. Space complexity is O(1) as no additional s
torage is needed.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func detectCycleV1[T constraints.Integer](
	head *internal.ListNode[T],
) *internal.ListNode[T] {
	if head == nil {
		return nil
	}

	getInersection := func() *internal.ListNode[T] {
		p1 := head
		p2 := head
		for (p2 != nil) && (p2.Next != nil) {
			p1 = p1.Next
			p2 = p2.Next.Next
			if p1 == p2 {
				return p1
			}
		}

		return nil
	}

	p1 := head
	p2 := getInersection()
	if p2 == nil {
		return nil
	}

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
