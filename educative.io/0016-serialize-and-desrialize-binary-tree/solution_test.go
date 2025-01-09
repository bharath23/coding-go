package educative0016

import (
	"testing"

	"github.com/bharath23/coding-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
	}{
		{
			name:  "test 1",
			input: [][]int{{100}, {50}, {200}, {25}, {75}, {350}},
		},
		{
			name:  "test 2",
			input: [][]int{{100}, {75}, {200}, {50}, {350}, {25}},
		},
		{
			name:  "test 3",
			input: [][]int{{1}, nil, {2}, nil, {3}, nil, {4}, nil, {5}},
		},
		{
			name:  "test 4",
			input: [][]int{{100}},
		},
		{
			name:  "test 5",
			input: [][]int{},
		},
		{
			name: "test 6",
			input: [][]int{{-6}, {572}, {653}, {-192}, {785}, {-435}, {-419},
				{745}, {613}, {-441}, {631}, {112}, {73}, {-252}, {-928},
				{-903}, {-994}, {-477}, {-636}, {-629}, {-886}, {-392},
				{-375}, {-37}, {-172}, {273}, {822}, {-66}, {589}, {-330},
				{696}, {463}, {-862}, {349}, nil, nil, {-994}, {-668}, nil,
				{296}, {-40}, {91}, {-584}, {206}, {-495}, {252}, {470},
				{-27}, {640}, {340}, {-349}, nil, nil, nil, nil, {-910},
				{623}, {-611}, {-361}, {-80}, {-757}, {445}, {526}, {-164},
				{-202}, {-522}, {-778}, nil, nil, nil, nil, nil, nil, {-847},
				nil, {774}, nil, nil, nil, {-471}, {-708}, nil, {-635},
				{-670}, {439}, {275}, nil, nil, {-116}, nil, nil, nil, nil,
				nil, nil, {745}, {-680}, nil, {-6}, nil, nil, nil, nil, nil,
				{946}, {351}, nil, nil, {-476}, nil, {-922}, nil, {831}, nil,
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
				nil, nil, {259}, nil, {198}, nil, nil, nil, nil, {989}, nil,
				nil, nil, nil, {115}, nil, nil, nil, nil, nil, nil, nil, nil,
				nil, {464}},
		},
	}

	for _, test := range tests {
		root := internal.NewTreeFromList[int](test.input)
		serialized := serialize[int](root)
		deserialized := deserialize[int](&serialized)
		have := internal.NewListFromTree[int](deserialized)
		assert.Equalf(t, test.input, have,
			"%s: serialization/deserialization failed", test.name)
	}
}
