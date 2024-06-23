package leetcode0206

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
		head: []int{1, 2, 3, 4, 5},
		want: []int{5, 4, 3, 2, 1},
	},
	{
		name: "test 2",
		head: []int{1, 2},
		want: []int{2, 1},
	},
	{
		name: "test 3",
		head: []int{},
		want: []int{},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := reverseListV0(head)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: reverse list does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := reverseListV1(head)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: reverse list does not match", test.name)
	}
}
