package leetcode0328

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
		want: []int{1, 3, 5, 2, 4},
	},
	{
		name: "test 2",
		head: []int{2, 1, 3, 5, 6, 4, 7},
		want: []int{2, 3, 6, 7, 1, 5, 4},
	},
	{
		name: "test 2",
		head: []int{1, 2, 3, 4, 5, 6, 7, 8},
		want: []int{1, 3, 5, 7, 2, 4, 6, 8},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := oddEvenListV0(head)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list do not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := oddEvenListV1(head)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list do not match", test.name)
	}
}
