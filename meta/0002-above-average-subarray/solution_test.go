package meta0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  [][2]int
	}{
		{
			name:  "test1",
			array: []int{3, 4, 2},
			want:  [][2]int{[2]int{1, 2}, [2]int{1, 3}, [2]int{2, 2}},
		},
	}

	for _, test := range tests {
		have := aboveAverageSubarrays(test.array)
		assert.Equalf(t, test.want, have, "%s: aboveAverageSubarrays mismatch", test.name)
	}
}
