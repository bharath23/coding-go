package educative0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test1",
			s:    "g",
			want: "g",
		},
		{
			name: "test2",
			s:    "ggaabcdeb",
			want: "bcdeb",
		},
		/*
			{
				name: "test3",
				s:    "abbddaccaaabcd",
				want: "abcd",
			},
			{
				name: "test4",
				s:    "aabbccdd",
				want: "",
			},
			{
				name: "test5",
				s:    "aannkwwwkkkwna",
				want: "kwkwna",
			},
		*/
	}

	for _, test := range tests {
		have := removeDuplicatesV0(test.s)
		assert.Equalf(t, test.want, have, "%s: removeDuplicate incorrect result", test.name)
	}

	for _, test := range tests {
		str := make([]byte, len(test.s))
		copy(str, test.s)
		have := removeDuplicatesV1(string(str))
		assert.Equalf(t, test.want, have, "%s: removeDuplicate incorrect result", test.name)
	}
}
