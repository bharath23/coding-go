package leetcode0680

/*
 * The solution uses two pointers to compare characters from both ends.
 * When a mismatch occurs, it checks at most two substrings by skipping
 * one character from either side and verifying if the remainder is a
 * palindrome.
 *
 * In the worst case, each palindrome check takes O(n), and since it happens at
 * most once, the overall time complexity remains O(n).
 *
 * The function does not allocate additional data structures proportional to
 * input size, so space complexity is O(1).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(1)
 */
func validPalindrome(s string) bool {
	isPalindrome := func(str string, i, j int) bool {
		for i < j {
			if str[i] != str[j] {
				return false
			}
			i++
			j--
		}

		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
		}

		i++
		j--
	}

	return true
}
