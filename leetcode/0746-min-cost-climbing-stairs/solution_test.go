package leetcode0746

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "test 1",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "test 2",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
		{
			name: "test 3",
			cost: []int{1, 2},
			want: 1,
		},
	}

	for _, test := range tests {
		have := minCostClimbingStairsV0(test.cost)
		assert.Equalf(t, test.want, have, "%s: minCostClimbingStairsV0 results incorrect", test.name)
	}

	for _, test := range tests {
		have := minCostClimbingStairsV1(test.cost)
		assert.Equalf(t, test.want, have, "%s: minCostClimbingStairsV1 results incorrect", test.name)
	}

	for _, test := range tests {
		have := minCostClimbingStairsV2(test.cost)
		assert.Equalf(t, test.want, have, "%s: minCostClimbingStairsV2 results incorrect", test.name)
	}
}
