package leetcode0160

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name         string
	intersectVal int
	listA        []int
	listB        []int
	skipA        int
	skipB        int
}{
	{
		name:         "test 1",
		intersectVal: 8,
		listA:        []int{4, 1, 8, 4, 5},
		listB:        []int{5, 6, 1, 8, 4, 5},
		skipA:        2,
		skipB:        3,
	},
	{
		name:         "test 2",
		intersectVal: 2,
		listA:        []int{1, 9, 1, 2, 4},
		listB:        []int{3, 2, 4},
		skipA:        3,
		skipB:        1,
	},
	{
		name:         "test 3",
		intersectVal: 0,
		listA:        []int{2, 6, 4},
		listB:        []int{1, 5},
		skipA:        3,
		skipB:        2,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := internal.NewListsWithIntersection(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV0(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := internal.NewListsWithIntersection(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV1(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := internal.NewListsWithIntersection(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV2(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}
