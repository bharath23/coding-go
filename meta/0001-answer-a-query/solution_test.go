package meta0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		queries [][2]int
		want    []int
	}{
		{
			name:    "test 1",
			n:       5,
			queries: [][2]int{[2]int{2, 3}, [2]int{1, 2}, [2]int{2, 1}, [2]int{2, 3}, [2]int{2, 2}},
			want:    []int{-1, 2, -1, 2},
		},
	}

	for _, test := range tests {
		have := answerQueries(test.queries, test.n)
		assert.Equalf(t, test.want, have, "%s: answerQueries mismatch", test.name)
	}
}
