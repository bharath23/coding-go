package educative0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		want        bool
	}{
		{
			name:        "test1",
			inputString: "ABCDABCD",
			want:        false,
		},
		{
			name:        "test2",
			inputString: "hello",
			want:        false,
		},
		{
			name:        "test3",
			inputString: "RaCEACAR",
			want:        false,
		},
		{
			name:        "test4",
			inputString: "A",
			want:        true,
		},
		{
			name:        "test5",
			inputString: "kaYak",
			want:        true,
		},
	}

	for _, test := range tests {
		have := isPalindrome(test.inputString)
		assert.Equalf(t, test.want, have, "%s: isPalindrome incorrect result", test.name)
	}
}
