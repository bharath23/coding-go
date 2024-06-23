package leetcode0876

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "test 1",
			head: []int{1, 2, 3, 4, 5},
			want: []int{3, 4, 5},
		},
		{
			name: "test 2",
			head: []int{1, 2, 3, 4, 5, 6},
			want: []int{4, 5, 6},
		},
	}

	for _, test := range tests {
		head := internal.NewListFromSlice(test.head)
		have := internal.ListToSlice(middleNode(head))
		assert.Equalf(t, test.want, have, "%s: middleNode did not return expected node", test.name)
	}
}
