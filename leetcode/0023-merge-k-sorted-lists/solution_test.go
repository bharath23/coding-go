package leetcode0023

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func listToSlice(l *ListNode) []int {
	slice := []int{}
	cur := l
	for cur != nil {
		slice = append(slice, cur.val)
		cur = cur.next
	}

	return slice
}

func sliceToLists(slice [][]int) []*ListNode {
	lists := []*ListNode{}
	for _, s := range slice {
		var head, prev *ListNode
		for i, v := range s {
			n := &ListNode{val: v, next: nil}
			if i == 0 {
				head = n
			}

			if prev != nil {
				prev.next = n
			}

			prev = n
		}
		lists = append(lists, head)
	}

	return lists
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "test 1",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "test 2",
			lists: [][]int{},
			want:  []int{},
		},
		{
			name:  "test 3",
			lists: [][]int{{}},
			want:  []int{},
		},
		{
			name:  "test 4",
			lists: [][]int{{}, {1}},
			want:  []int{1},
		},
	}

	for _, test := range tests {
		lists := sliceToLists(test.lists)
		have := mergeKLists(lists)
		assert.Equalf(t, test.want, listToSlice(have), "%s: mergeKLists failed", test.name)
	}
}
