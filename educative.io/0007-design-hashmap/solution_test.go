package educative0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutiuon(t *testing.T) {
	var tests = []struct {
		name   string
		input1 []string
		input2 [][]int
		want   [][]int
	}{
		{
			name:   "test1",
			input1: []string{"DesignHashMap", "Put", "Get"},
			input2: [][]int{nil, {15, 250}, {15}},
			want:   [][]int{nil, nil, {250}},
		},
		{
			name:   "test2",
			input1: []string{"DesignHashMap", "Get", "Get", "Get", "Remove", "Get"},
			input2: [][]int{nil, {5}, {15}, {11}, {11}, {11}},
			want:   [][]int{nil, {-1}, {-1}, {-1}, nil, {-1}},
		},
		{
			name:   "test3",
			input1: []string{"DesignHashMap", "Get", "Get", "Put", "Get"},
			input2: [][]int{nil, {5}, {15}, {15, 250}, {15}},
			want:   [][]int{nil, {-1}, {-1}, nil, {250}},
		},
		{
			name:   "test4",
			input1: []string{"DesignHashMap", "Put", "Put", "Get", "Remove", "Get"},
			input2: [][]int{nil, {51, 300}, {15, 750}, {51}, {51}, {51}},
			want:   [][]int{nil, nil, nil, {300}, nil, {-1}},
		},
		{
			name:   "test5",
			input1: []string{"DesignHashMap", "Put", "Put", "Get", "Remove", "Get"},
			input2: [][]int{nil, {51, 300}, {15, 750}, {51}, {51}, {50}},
			want:   [][]int{nil, nil, nil, {300}, nil, {-1}},
		},
		{
			name:   "test6",
			input1: []string{"DesignHashMap", "Put", "Get", "Put", "Get", "Remove", "Get"},
			input2: [][]int{nil, {51, 300}, {51}, {51, 750}, {51}, {51}, {51}},
			want:   [][]int{nil, nil, {300}, nil, {750}, nil, {-1}},
		},
	}

	for _, test := range tests {
		var hashMap *DesignHashMap
		have := [][]int{}
		for i, inp1 := range test.input1 {
			switch inp1 {
			case "DesignHashMap":
				hm := Constructor()
				hashMap = &hm
				have = append(have, nil)
			case "Get":
				g := hashMap.Get(test.input2[i][0])
				have = append(have, []int{g})
			case "Put":
				hashMap.Put(test.input2[i][0], test.input2[i][1])
				have = append(have, nil)
			case "Remove":
				hashMap.Remove(test.input2[i][0])
				have = append(have, nil)
			}
		}
		assert.Equalf(t, test.want, have, "%s: incorrect DesignHashMap implementation", test, test.name)
	}
}
