package educative0020

import (
	"testing"

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
			input: []int{1, 2, 2, 3, 3, 4},
			want:  []int{1, 4},
		},
		{
			name:  "test 2",
			input: []int{2, 1, 1, 8},
			want:  []int{2, 8},
		},
		{
			name:  "test 3",
			input: []int{4, 1, 2, 1, 2, 0},
			want:  []int{4, 0},
		},
		{
			name:  "test 4",
			input: []int{2, 5},
			want:  []int{2, 5},
		},
		{
			name:  "test 5",
			input: []int{4, 5, 4, 1, 1, 2, 6, 6},
			want:  []int{5, 2},
		},
		{
			name: "test 6",
			input: []int{-2147483648, 2147483647, 2147483647, -2147483648,
				914124654, 1708175908},
			want: []int{914124654, 1708175908},
		},
	}

	for _, test := range tests {
		have := twoSingleNumbers(test.input)
		assert.ElementsMatchf(t, test.want, have, "%s: twoSingleNumbers "+
			"incorrect", test.name)
	}
}
