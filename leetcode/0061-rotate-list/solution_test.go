package leetcode0061

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	have []int
	k    int
	want []int
}{
	{
		name: "test 1",
		have: []int{1, 2, 3, 4, 5},
		k:    2,
		want: []int{4, 5, 1, 2, 3},
	},
	{
		name: "test 2",
		have: []int{0, 1, 2},
		k:    4,
		want: []int{2, 0, 1},
	},
	{
		name: "test 3",
		have: []int{1, 2},
		k:    0,
		want: []int{1, 2},
	},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.have)
		haveList := rotateRight(head, test.k)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: rotated list do not match", test.name)
	}
}
