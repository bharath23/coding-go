package educative0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  int
	}{
		{
			name:  "test 1",
			array: []int{0, 1, 2, 4},
			want:  3,
		},
		{
			name:  "test 2",
			array: []int{3, 0, 1, 4},
			want:  2,
		},
		{
			name:  "test 3",
			array: []int{1, 4, 5, 6, 8, 2, 0, 7},
			want:  3,
		},
		{
			name:  "test 4",
			array: []int{1, 0, 2, 3, 4, 5, 6, 8, 9, 7, 11},
			want:  10,
		},
		{
			name:  "test 5",
			array: []int{1},
			want:  0,
		},
		{
			name:  "test 6",
			array: []int{0},
			want:  1,
		},
	}

	for _, test := range tests {
		have := findMissingNumber(test.array)
		assert.Equalf(t, test.want, have,
			"%s: findMissingNumber returned incorrect value", test.name)
	}
}
