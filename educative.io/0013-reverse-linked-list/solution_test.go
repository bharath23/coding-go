package educative0013

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test 1",
			input: []int{1, -2, 3, 4, -5, 4, 3, -2, 1},
			want:  []int{1, -2, 3, 4, -5, 4, 3, -2, 1},
		},
		{
			name:  "test 2",
			input: []int{-1, -5, -3, -7, -8, -6, -2},
			want:  []int{-2, -6, -8, -7, -3, -5, -1},
		},
		{
			name:  "test 3",
			input: []int{-1, 2, -3, 4},
			want:  []int{4, -3, 2, -1},
		}, {
			name:  "test 4",
			input: []int{1, -1, -2, 3, -4, 5},
			want:  []int{5, -4, 3, -2, -1, 1},
		}, {
			name:  "test 5",
			input: []int{28, 21, 14, 7},
			want:  []int{7, 14, 21, 28},
		},
	}

	for _, test := range tests {
		head := internal.NewListFromSlice(test.input)
		haveList := reverse(head)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list do not match ", test.name)
	}
}
