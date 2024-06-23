package leetcode0234

import (
	"github.com/bharath23/coding-go/internal"
)

/*
Simple solution which uses copy of the data to check if the linked list is
palindrome. Time complexity is O(n) to make the copy of the values, another
O(n) to make a copy of values in reverse and O(n) to compare the values.
Total time complexity is 3*O(n) or O(n). Space complexity is O(n) for copy
of the values and another O(n) for the values in reverse. Total space
complexity is 2*O(n) or O(n).

Complexity:
Time complexity: O(n)
Space complexity: O(n)
*/

func isPalindromeV0(head *internal.ListNode) bool {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}

	reverseVals := []int{}
	for i := len(vals) - 1; i >= 0; i-- {
		reverseVals = append(reverseVals, vals[i])
	}

	for i := 0; i < len(vals); i++ {
		if vals[i] != reverseVals[i] {
			return false
		}
	}

	return true
}

/*
Recursive solution. We traverse the list once. Time complexity is O(n), we
traverse the list once. The space complexity is O(1) as no additional storage
is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/

func isPalindromeV1(head *internal.ListNode) bool {
	var cur *internal.ListNode
	var recursiveCheck func(*internal.ListNode) bool
	recursiveCheck = func(node *internal.ListNode) bool {
		if node != nil {
			if !recursiveCheck(node.Next) {
				return false
			}
			if cur.Val != node.Val {
				return false
			}
			cur = cur.Next
		}

		return true
	}

	cur = head
	return recursiveCheck(head)
}

/*
We reverse the second half of the list and compare the values to check if the
list is a palindrome. Time complexity is O(n) to find the end of the first half
O(n) to reverse the second half of the list. O(n) to compare the values. O(n)
to reverse the second half. Total time complexity is 4*O(n) or O(n). Space
complexity is O(1) as no additional storage is required.

Complexity:
Time complexity: O(n)
Space complexity: O(1)
*/
func isPalindromeV2(head *internal.ListNode) bool {
	findFirstHalfEnd := func(head *internal.ListNode) *internal.ListNode {
		fast := head
		slow := head
		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}

		return slow
	}

	reverseList := func(head *internal.ListNode) *internal.ListNode {
		var prev *internal.ListNode
		cur := head
		for cur != nil {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}

		return prev
	}

	firstHalfEnd := findFirstHalfEnd(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)
	result := true
	p1 := head
	p2 := secondHalfStart
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)

	return result
}
