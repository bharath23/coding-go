package educative0025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			name:  "test 1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{4, 5, 6},
			n:     3,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "test 2",
			nums1: []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0},
			m:     5,
			nums2: []int{-1, -1, 0, 0, 1, 2},
			n:     6,
			want:  []int{-1, -1, -1, 0, 0, 0, 0, 0, 1, 2, 3},
		}, {
			name:  "test 3",
			nums1: []int{6, 7, 8, 9, 10, 0, 0, 0, 0, 0},
			m:     5,
			nums2: []int{1, 2, 3, 4, 5},
			n:     5,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			name:  "test 4",
			nums1: []int{10, 49, 99, 110, 176, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			m:     5,
			nums2: []int{1, 2, 4, 7, 8, 12, 15, 19, 24, 50, 69, 80, 100},
			n:     13,
			want:  []int{1, 2, 4, 7, 8, 10, 12, 15, 19, 24, 49, 50, 69, 80, 99, 100, 110, 176},
		}, {
			name:  "test 5",
			nums1: []int{0, 1, 4, 9, 0, 0, 0, 0, 0, 0},
			m:     4,
			nums2: []int{-111, -20, -5, 5, 11, 20},
			n:     6,
			want:  []int{-111, -20, -5, 0, 1, 4, 5, 9, 11, 20},
		},
	}

	for _, test := range tests {
		have := mergeSorted(test.nums1, test.m, test.nums2, test.n)
		assert.Equalf(t, test.want, have, "%s: mergeSorted incorrect merge",
			test.name)
	}
}
