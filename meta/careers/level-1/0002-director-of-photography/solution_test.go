package meta0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int32
		c    string
		x    int32
		y    int32
		want int32
	}{
		{
			name: "test 1",
			n:    5,
			c:    "APABA",
			x:    1,
			y:    2,
			want: 1,
		},
		{
			name: "test 2",
			n:    5,
			c:    "APABA",
			x:    2,
			y:    3,
			want: 0,
		},
		{
			name: "test 3",
			n:    8,
			c:    ".PBAAP.B",
			x:    1,
			y:    3,
			want: 3,
		},
	}

	for _, test := range tests {
		have := getArtisticPhotographCount(test.n, test.c, test.x, test.y)
		assert.Equalf(t, test.want, have, "%s: getArtisticPhotographCount failed",
			test.name)
	}
}
