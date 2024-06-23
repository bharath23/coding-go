package leetcode0708

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name      string
	head      []int
	insertVal int
	want      []int
}{
	{
		name:      "test 1",
		head:      []int{3, 4, 1},
		insertVal: 2,
		want:      []int{3, 4, 1, 2},
	},
	{
		name:      "test 2",
		head:      []int{},
		insertVal: 1,
		want:      []int{1},
	},
	{
		name:      "test 3",
		head:      []int{1},
		insertVal: 0,
		want:      []int{1, 0},
	},
	{
		name:      "test 4",
		head:      []int{1, 3, 5},
		insertVal: 4,
		want:      []int{1, 3, 4, 5},
	},
	{
		name:      "test 5",
		head:      []int{3, 5, 1},
		insertVal: 0,
		want:      []int{3, 5, 0, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		head := internal.NewCircularListFromSlice(test.head)
		haveList := insert(head, test.insertVal)
		have := internal.CircularListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: circular list mismatch", test.name)
	}
}
