package leetcode0071

import (
	"strings"
)

/*
 * Splitting the string and iterating over all the segments requires O(n) time,
 * where n is the lenght of the input path. Each segment is processed in once,
 * and stack operations are O(1). Joining the final stack also takes O(n) time.
 *
 * Space complexity is also O(n), when all segments are stored in the stack.
 *
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func simplifyPath(path string) string {
	segments := strings.Split(path, "/")
	stack := []string{}
	for _, s := range segments {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s != "." && len(s) > 0 {
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
