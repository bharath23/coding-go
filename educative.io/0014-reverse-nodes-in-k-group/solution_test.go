package educative0014

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		k     int
		want  []int
	}{
		{
			name:  "test 1",
			input: []int{1, 2, 3, 4, 5},
			k:     2,
			want:  []int{2, 1, 4, 3, 5},
		},
		{
			name:  "test 2",
			input: []int{3, 4, 5, 6, 2, 8, 7, 7},
			k:     3,
			want:  []int{5, 4, 3, 8, 2, 6, 7, 7},
		},
		{
			name:  "test 3",
			input: []int{1, 2, 3, 4, 5},
			k:     1,
			want:  []int{1, 2, 3, 4, 5},
		}, {
			name:  "test 4",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     3,
			want:  []int{3, 2, 1, 6, 5, 4, 7},
		}, {
			name:  "test 5",
			input: []int{1, 2, 3, 4, 5, 6, 7},
			k:     7,
			want:  []int{7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		head := internal.NewListFromSlice(test.input)
		haveList := reverseKGroups(head, test.k)
		have := internal.ListToSlice(haveList)
		assert.Equalf(t, test.want, have, "%s: list do not match ", test.name)
	}
}
