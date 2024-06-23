package leetcode0203

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	val  int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 6, 3, 4, 5, 6},
		val:  6,
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "test 2",
		head: []int{},
		val:  1,
		want: []int{},
	},
	{
		name: "test 1",
		head: []int{7, 7, 7, 7},
		val:  7,
		want: []int{},
	},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		haveList := removeElements(head, test.val)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: post deletion list mismatch", test.name)
	}
}
