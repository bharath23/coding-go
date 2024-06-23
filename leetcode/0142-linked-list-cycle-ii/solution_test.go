package leetcode0141

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	input []int
	pos   int
}{
	{
		name:  "test 1",
		input: []int{3, 2, 0, -4},
		pos:   1,
	},
	{
		name:  "test 2",
		input: []int{1, 2},
		pos:   0,
	},
	{
		name:  "test 3",
		input: []int{1},
		pos:   -1,
	},
	{
		name:  "test 4",
		input: []int{},
		pos:   -1,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head, want := internal.NewListFromSliceWithCycle(test.input, test.pos)
		have := detectCycleV0(head)
		assert.Equalf(t, want, have, "%s: testing for cycle mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head, want := internal.NewListFromSliceWithCycle(test.input, test.pos)
		have := detectCycleV1(head)
		assert.Equalf(t, want, have, "%s: testing for cycle mismatch", test.name)
	}
}
