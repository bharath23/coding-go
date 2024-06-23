package leetcode0138

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head [][]int
	want [][]int
}{
	{
		name: "test 1",
		head: [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		want: [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
	},
	{
		name: "test 2",
		head: [][]int{{1, 1}, {2, 1}},
		want: [][]int{{1, 1}, {2, 1}},
	},
	{
		name: "test 3",
		head: [][]int{{3, -1}, {3, 0}, {3, -1}},
		want: [][]int{{3, -1}, {3, 0}, {3, -1}},
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewRandomListFromSlice(test.head)
		haveList := copyRandomListV0(head)
		have := internal.RandomListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: copied list does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewRandomListFromSlice(test.head)
		haveList := copyRandomListV1(head)
		have := internal.RandomListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: copied list does not match", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		head := internal.NewRandomListFromSlice(test.head)
		haveList := copyRandomListV2(head)
		have := internal.RandomListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: copied list does not match", test.name)
	}
}

func TestSolutionV3(t *testing.T) {
	for _, test := range tests {
		head := internal.NewRandomListFromSlice(test.head)
		haveList := copyRandomListV3(head)
		have := internal.RandomListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: copied list does not match", test.name)
	}
}
