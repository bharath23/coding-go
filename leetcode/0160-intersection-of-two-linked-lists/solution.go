package leetcode0160

import (
	"github.com/bharath23/coding-go/internal"
	"golang.org/x/exp/constraints"
)

/*
Simple two pass solution. For each node in listA compare with each node in
listB. If they are equal return the node. Time complexity is o(n^2), worst
case you have to visit each node. Space complexity is O(1) as there is no
additional space requrired.

Complexity:
Time complexity: O(n^2)
Space complexity: O(1)
*/

func getIntersectionNodeV0[T constraints.Integer](
	headA, headB *internal.ListNode[T],
) *internal.ListNode[T] {
	curA := headA
	for curA != nil {
		curB := headB
		for curB != nil {
			if curA == curB {
				return curA
			}
			curB = curB.Next
		}
		curA = curA.Next
	}

	return nil
}

/*
Simple one pass soluation using hash map. Visit each node of listA and add it
to a visited hash map. Visit each node of listB and check if it has already
been visited. Time complexity is O(n) as each node is just visited once. Space
complexity is O(n), required to store the hashmap.

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func getIntersectionNodeV1[T constraints.Integer](
	headA, headB *internal.ListNode[T],
) *internal.ListNode[T] {
	visited := map[*internal.ListNode[T]]bool{}
	curA := headA
	for curA != nil {
		visited[curA] = true
		curA = curA.Next
	}

	curB := headB
	for curB != nil {
		isVisited := visited[curB]
		if isVisited {
			return curB
		}

		curB = curB.Next
	}

	return nil
}

/*
A two pointer solution. Two pointers traverse the lists. When one reaches the
end of the list it restarts at the start of the second list. During the second
iteration the pointers will meet at the intersection node. Time complexity is
O(n) and space (1).

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func getIntersectionNodeV2[T constraints.Integer](
	headA, headB *internal.ListNode[T],
) *internal.ListNode[T] {
	ptrA := headA
	ptrB := headB
	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}

		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}

	return ptrA
}
