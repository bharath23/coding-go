package leetcode0071

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "test 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "test 2",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "test 3",
			path: "/home/user/Documents/../Pictures",
			want: "/home/user/Pictures",
		},
		{
			name: "test 4",
			path: "/../",
			want: "/",
		},
		{
			name: "test 5",
			path: "/.../a/../b/c/../d/./",
			want: "/.../b/d",
		},
	}

	for _, test := range tests {
		have := simplifyPath(test.path)
		assert.Equalf(t, test.want, have, "%s: simplifyPath failed", test.name)
	}
}
