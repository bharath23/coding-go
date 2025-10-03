package leetcode1249

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "test 1",
			s:    "lee(t(c)o)de)",
			want: []string{"lee(t(c)o)de", "lee(t(co)de)", "lee(t(c)ode)"},
		},
		{
			name: "test 2",
			s:    "a)b(c)d",
			want: []string{"ab(c)d"},
		},
		{
			name: "test 1",
			s:    "))((",
			want: []string{},
		},
	}

	for _, test := range tests {
		have := minRemoveToMakeValidv0(test.s)
		if have == "" {
			assert.Emptyf(t, test.want, "%s, minRemoveToMakeValidv0 failed", test.name)
		} else {
			assert.Containsf(t, test.want, have, "%s: minRemoveToMakeValidv0 failed", test.name)
		}
		have = minRemoveToMakeValidv1(test.s)
		if have == "" {
			assert.Emptyf(t, test.want, "%s, minRemoveToMakeValidv1 failed", test.name)
		} else {
			assert.Containsf(t, test.want, have, "%s: minRemoveToMakeValidv1 failed", test.name)
		}
		have = minRemoveToMakeValidv2(test.s)
		if have == "" {
			assert.Emptyf(t, test.want, "%s, minRemoveToMakeValidv2 failed", test.name)
		} else {
			assert.Containsf(t, test.want, have, "%s: minRemoveToMakeValidv2 failed", test.name)
		}
	}
}
