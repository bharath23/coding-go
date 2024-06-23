package leetcode0430

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 3, 4, 5, 6, -1, -1, -1, 7, 8, 9, 10, -1, -1, 11, 12},
		want: []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
	},
	{
		name: "test 2",
		head: []int{1, 2, -1, 3},
		want: []int{1, 3, 2},
	},
	{
		name: "test 3",
		head: []int{},
		want: []int{},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewMultiLevelDoublyLinkedList(test.head)
		haveList := flattenV0(head)
		have := internal.MultiLevelDoublyLinkedListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: flattened list do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewMultiLevelDoublyLinkedList(test.head)
		haveList := flattenV1(head)
		have := internal.MultiLevelDoublyLinkedListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: flattened list do not match", test.name)
	}
}
