package edutcative0017

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  bool
	}{
		{
			name:  "test 1",
			input: [][]int{{1}, {2}, {2}, {3}, {4}, {4}, {3}},
			want:  true,
		},
		{
			name:  "test 2",
			input: [][]int{{18}, {21}, {21}, {47}, {20}, {21}, {47}},
			want:  false,
		},
		{
			name:  "test 3",
			input: [][]int{{25}, {4}, {67}, {2}, {3}, {3}, {2}},
			want:  false,
		},
		{
			name:  "test 4",
			input: [][]int{{1}, {2}, {2}, {3}, nil, nil, {3}},
			want:  true,
		},
		{
			name:  "test 5",
			input: [][]int{{1}, {2}, {2}, nil, {3}, nil, {3}},
			want:  false,
		},
	}

	for _, test := range tests {
		root := internal.NewTreeFromList[int](test.input)
		have := isSymmetric(root)
		assert.Equalf(t, test.want, have, "%s: isSymmetric check failed",
			test.name)
	}
}
