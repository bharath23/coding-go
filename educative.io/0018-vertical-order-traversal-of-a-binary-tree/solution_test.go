package educative0018

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "test 1",
			input: [][]int{{100}, {50}, {200}, {25}, {75}, {300}, {10}, {350},
				{15}},
			want: [][]int{{350}, {25}, {50, 15}, {100, 75, 300}, {200}, {10}},
		},
		{
			name:  "test 2",
			input: [][]int{{20}, {40}, {50}, {90}, {67}, {94}},
			want:  [][]int{{90}, {40}, {20, 67, 94}, {50}},
		}, {
			name:  "test 3",
			input: [][]int{{10}, {23}, {45}, {25}, {46}},
			want:  [][]int{{25}, {23}, {10, 46}, {45}},
		}, {
			name:  "test 4",
			input: [][]int{{9}, {7}, nil, nil, {1}, {8}, {10}, nil, {12}},
			want:  [][]int{{7, 8}, {9, 1, 12}, {10}},
		}, {
			name:  "test 5",
			input: [][]int{{3}, {2}, {3}, nil, {3}, nil, {1}},
			want:  [][]int{{2}, {3, 3}, {3}, {1}},
		},
	}

	for _, test := range tests {
		root := internal.NewTreeFromList[int](test.input)
		have := verticalOrder(root)
		assert.Equalf(t, test.want, have, "%s: vertical orders do not match",
			test.name)
	}
}
