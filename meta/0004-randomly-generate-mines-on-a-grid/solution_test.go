package meta0004

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		k    int
	}{
		{
			name: "test 1",
			m:    5,
			n:    5,
			k:    9,
		},
	}

	for _, test := range tests {
		generateMinesOnGrid(test.m, test.n, test.k)
	}
}
