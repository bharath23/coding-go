package leetcode0864

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want int
	}{
		{
			name: "test 1",
			grid: []string{"@.a..", "###.#", "b.A.B"},
			want: 8,
		},
		{
			name: "test 2",
			grid: []string{"@..aA", "..B#.", "....b"},
			want: 6,
		},
		{
			name: "test 3",
			grid: []string{"@Aa"},
			want: -1,
		},
	}

	for _, test := range tests {
		have := shortestPathAllKeys(test.grid)
		assert.Equalf(t, test.want, have, "%s: shortestPathAllKeys failed", test.name)

	}
}
