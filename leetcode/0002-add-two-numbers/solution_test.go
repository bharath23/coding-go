package leetcode0002

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	l1   []int
	l2   []int
	want []int
}{
	{
		name: "test 1",
		l1:   []int{2, 4, 3},
		l2:   []int{5, 6, 4},
		want: []int{7, 0, 8},
	},
	{
		name: "test 2",
		l1:   []int{5},
		l2:   []int{5},
		want: []int{0, 1},
	},
	{
		name: "test 3",
		l1:   []int{0},
		l2:   []int{0},
		want: []int{0},
	},
	{
		name: "test 4",
		l1:   []int{9, 9, 9, 9, 9, 9, 9},
		l2:   []int{9, 9, 9, 9},
		want: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		l1 := internal.NewListFromSlice(test.l1)
		l2 := internal.NewListFromSlice(test.l2)
		haveList := addTwoNumbers(l1, l2)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: sum do not match", test.name)
	}
}
