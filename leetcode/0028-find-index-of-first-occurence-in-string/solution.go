package leetcode0028

import "fmt"

/*
 * The outer loop runs (n - m + 1) times in the worst case. For each index i,
 * the substring comparison haystack[i:i+m] == needle takes O(m) time because
 * string comparison checks each character. Overall time complexity is
 * O(n × m)
 * The space complexity is O(1) since no additional space is required.
 *
 * Time complexity: O(nxm)
 * Space complexity: O(1)
 */

func strStrv0(haystack string, needle string) int {
	m := len(needle)
	n := len(haystack)
	if m > n {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

/*
* This approach uses a straightforward sliding window technique. Each time we
* move the window, comparing the substring with the target takes O(m) time. As
* the window slides across the string, the overall time complexity becomes
* O(n × m).
* The space complexity is O(1) since no additional space is required.
*
* Time complexity: O(nxm)
* Space complexity: O(1)
 */

func strStrv1(haystack string, needle string) int {
	m := len(needle)
	n := len(haystack)
	if m > n {
		return -1
	}

	for window_start := 0; window_start <= (n - m); window_start++ {
		for i := 0; i < m; i++ {
			fmt.Printf("needle[%d]: %c, haystack[%d]: %c\n", i, needle[i],
				window_start+i, haystack[window_start+i])
			if needle[i] != haystack[window_start+i] {
				break
			}

			if i == m-1 {
				return window_start
			}
		}
	}

	return -1
}

func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	lps[0] = 0
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[i-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

/*
* This approach uses the Knuth-Morris-Pratt (KMP) algorithm, which preprocesses
* the pattern to build an LPS (Longest Prefix Suffix) array. This allows the
* search to skip rechecking characters, avoiding redundant comparisons.
*
* Time complexity: O(n + m)
* Space complexity: O(m)
 */

func strStrKMP(haystack string, needle string) int {
	m := len(needle)
	n := len(haystack)
	if m > n {
		return -1
	}

	lps := buildLPS(needle)
	i := 0
	j := 0
	for i < n {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == m {
				return i - j
			}
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}
