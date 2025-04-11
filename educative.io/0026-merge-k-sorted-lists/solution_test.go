package educative0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "test 1",
			lists: [][]int{{2}, {2}, {1, 2, 4}},
			want:  []int{1, 2, 2, 2, 4},
		},
		{
			name:  "test 2",
			lists: [][]int{{11, 41, 51}, {21, 23, 42}},
			want:  []int{11, 21, 23, 41, 42, 51},
		},
		{
			name:  "test 3",
			lists: [][]int{{2}, {1, 2, 4}, {25, 56, 66, 72}},
			want:  []int{1, 2, 2, 4, 25, 56, 66, 72},
		},
		{
			name:  "test 4",
			lists: [][]int{{11, 41, 51}, {2}},
			want:  []int{2, 11, 41, 51},
		},
		{
			name:  "test 5",
			lists: [][]int{{21, 23, 42}, {1, 2, 4}},
			want:  []int{1, 2, 4, 21, 23, 42},
		},
	}

	for _, test := range tests {
		lists := slicesToLists(test.lists)
		have := mergeKLists(lists)
		assert.Equalf(t, test.want, listToSlice(have),
			"%s: mergeKLists merge failed", test.name)
	}
}

func slicesToLists(slices [][]int) []*LinkedList {
	result := []*LinkedList{}
	for _, s := range slices {
		var prev *LinkedListNode
		for i, e := range s {
			n := &LinkedListNode{val: e}
			if i == 0 {
				result = append(result, &LinkedList{head: n})
			}

			if prev != nil {
				prev.next = n
			}

			prev = n
		}
	}

	return result
}

func listToSlice(head *LinkedListNode) []int {
	result := []int{}
	cur := head
	for cur != nil {
		result = append(result, cur.val)
		cur = cur.next
	}

	return result
}
