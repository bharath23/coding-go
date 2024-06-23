package leetcode0019

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	n    int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 3, 4, 5},
		n:    2,
		want: []int{1, 2, 3, 5},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := removeNthFromEndV0(head, test.n)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list different", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := removeNthFromEndV0(head, test.n)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list different", test.name)
	}
}
