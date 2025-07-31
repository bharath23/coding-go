package leetcode0025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func listToSlice(head *ListNode) []int {
	slice := []int{}
	cur := head
	for cur != nil {
		slice = append(slice, cur.Val)
		cur = cur.Next
	}

	return slice
}

func sliceToList(slice []int) *ListNode {
	var cur, head *ListNode
	for _, v := range slice {
		n := &ListNode{
			Val: v,
		}
		if head == nil {
			head = n
		}
		if cur != nil {
			cur.Next = n
		}
		cur = n
	}

	return head
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			name: "test 1",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "test 2",
			head: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},
	}

	for _, test := range tests {
		head := sliceToList(test.head)
		have := listToSlice(reverseKGroupv0(head, test.k))
		assert.Equalf(t, test.want, have, "%s: reverseKGroupv0 failed",
			test.name)
	}

	for _, test := range tests {
		head := sliceToList(test.head)
		have := listToSlice(reverseKGroupv1(head, test.k))
		assert.Equalf(t, test.want, have, "%s: reverseKGroupv1 failed",
			test.name)
	}
}
