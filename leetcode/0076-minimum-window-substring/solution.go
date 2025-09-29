package leetcode0076

/*
 * The solution uses a sliding window technique. Each character in s is
 * processed at most twice: once when the righ pointer expands and once when
 * the left pointer contacts the window. All map operations are O(1). The
 * overall time complexity is O(m+n), where m is the length of s and n is the
 * length of t.
 *
 * The space complexity is O(m+n), which tracks the number of unique characters,
 * across s and t.
 *
 * Time complexity: O(m+n)
 * Space complexity: O(m+n)
 *
 */
func minWindowv0(s, t string) string {
	m := len(s)
	n := len(t)
	if m == 0 || n == 0 {
		return ""
	}

	tCount := make(map[rune]int)
	for _, c := range t {
		tCount[c]++
	}

	required := len(tCount)
	l, r := 0, 0
	complete := 0
	windowCount := make(map[rune]int)
	res := []int{-1, 0, 0}
	for r < m {
		c := rune(s[r])
		windowCount[c]++
		if _, ok := tCount[c]; ok && tCount[c] == windowCount[c] {
			complete++
		}

		for l <= r && complete == required {
			c = rune(s[l])
			if res[0] == -1 || r-l+1 < res[0] {
				res[0] = r - l + 1
				res[1] = l
				res[2] = r
			}

			windowCount[c]--
			if _, ok := tCount[c]; ok && windowCount[c] < tCount[c] {
				complete--
			}
			l++
		}
		r++
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[1] : res[2]+1]
}

/*
 * In this solution we prefilter the string s to only have characters from the
 * string t. If the length of the filtered string is f, the time complexity is
 * O(f). Each character again is visited at most twice, once when the right
 * pointer expands the window and once when the left pointer contracts the
 * window. The total time complexity is O(f+n), which in the worst case is
 * O(m+n).
 * The solution uses a sliding window technique. Each character in s is
 * processed at most twice: once when the righ pointer expands and once when
 * the left pointer contacts the window. All map operations are O(1). The
 * overall time complexity is O(m+n), where m is the length of s and n is the
 * length of t.
 *
 * The space complexity is O(f+k), where k is the number of unique characters
 * in t. Worst case the space complexity is O(m+n).
 *
 * Time complexity: O(m+n)
 * Space complexity: O(m+n)
 *
 */
func minWindowv1(s, t string) string {
	m := len(s)
	n := len(t)
	if m == 0 || n == 0 {
		return ""
	}

	tCount := make(map[rune]int)
	for _, c := range t {
		tCount[c]++
	}

	required := len(tCount)
	filtered := []struct {
		element rune
		index   int
	}{}
	for i, c := range s {
		if _, ok := tCount[c]; ok {
			filtered = append(filtered,
				struct {
					element rune
					index   int
				}{
					element: c,
					index:   i,
				},
			)
		}
	}

	l, r := 0, 0
	complete := 0
	windowCount := make(map[rune]int)
	res := []int{-1, 0, 0}

	for r < len(filtered) {
		c := filtered[r].element
		windowCount[c]++
		if count, ok := tCount[c]; ok && windowCount[c] == count {
			complete++
		}

		for l <= r && complete == required {
			c := filtered[l].element
			s := filtered[l].index
			e := filtered[r].index
			if res[0] == -1 || e-s+1 < res[0] {
				res[0] = e - s + 1
				res[1] = s
				res[2] = e
			}
			windowCount[c]--
			if count, ok := tCount[c]; ok && windowCount[c] < count {
				complete--
			}
			l++
		}
		r++
	}
	if res[0] == -1 {
		return ""
	}

	return s[res[1] : res[2]+1]
}
