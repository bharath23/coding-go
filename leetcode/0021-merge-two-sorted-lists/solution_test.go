package leetcode0021

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	list1 []int
	list2 []int
	want  []int
}{
	{
		name:  "test 1",
		list1: []int{1, 2, 4},
		list2: []int{1, 3, 4},
		want:  []int{1, 1, 2, 3, 4, 4},
	},
	{
		name:  "test 2",
		list1: []int{},
		list2: []int{},
		want:  []int{},
	},
	{
		name:  "test 3",
		list1: []int{},
		list2: []int{0},
		want:  []int{0},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		list1 := internal.NewListFromSlice(test.list1)
		list2 := internal.NewListFromSlice(test.list2)
		haveList := mergeTwoListsV0(list1, list2)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: merged list dont match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		list1 := internal.NewListFromSlice(test.list1)
		list2 := internal.NewListFromSlice(test.list2)
		haveList := mergeTwoListsV1(list1, list2)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: merged list dont match", test.name)
	}
}
