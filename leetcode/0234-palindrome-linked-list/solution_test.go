package leetcode0234

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	want bool
}{
	{
		name: "test 1",
		head: []int{1, 2, 2, 1},
		want: true,
	},
	{
		name: "test 2",
		head: []int{1, 2},
		want: false,
	},
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		have := isPalindromeV0(head)
		assert.Equalf(t, test.want, have, "%s: isPalindrome check failed", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		have := isPalindromeV1(head)
		assert.Equalf(t, test.want, have, "%s: isPalindrome check failed", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		have := isPalindromeV2(head)
		assert.Equalf(t, test.want, have, "%s: isPalindrome check failed", test.name)
	}
}
