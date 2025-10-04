package leetcode0408

/*
 * The solution uses a single loop that iterates through both the word and its
 * abbreviation simultaneously. Each character in the abbreviation is processed
 * once. When digits are encountered, they are parsed as a complete number, and
 * the word index advances by that count. Therefore, both word and abbreviation
 * are traversed linearly, leading to an overall time complexity of O(n), where
 * n is the length of the abbreviation.
 *
 * The solution does use any additional data structures, hence the space
 * complexity is O(1).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(1)
 */
func validWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if abbr[j] >= '0' && abbr[j] <= '9' {
			if abbr[j] == '0' {
				return false
			}

			num := 0
			for j < len(abbr) && abbr[j] >= '0' && abbr[j] <= '9' {
				num = 10*num + int(abbr[j]-'0')
				j++
			}
			i += num
		} else {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		}
	}

	return i == len(word) && j == len(abbr)
}
