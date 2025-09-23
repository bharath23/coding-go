package leetcode0022

/*
 * The algorithm generates all valid combinations of n pairs of parentheses
 * using backtracking. At each step, it can add either a left parenthesis
 * (if fewer than n have been placed) or a right parenthesis (if fewer than
 * the number of left parentheses placed). This ensures that only valid
 * prefixes are explored.
 *
 * The total number of valid combinations corresponds to the nth Catalan
 * number, C_n, which grows asymptotically as 4^n / (n^(3/2)).
 * Since each valid string has length 2n, the time to build all results is
 * O(C_n * n) ≈ O(4^n / sqrt(n)).
 *
 * The space complexity is O(C_n * n) for storing the result list, plus
 * O(n) for the recursion stack.
 *
 * Time complexity: O(4^n / sqrt(n))
 * Space complexity: O(4^n / sqrt(n))
 */
func generateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(curString string, left, right int)

	backtrack = func(curString string, left int, right int) {
		if len(curString) == 2*n {
			result = append(result, curString)
		}

		if left < n {
			backtrack(curString+"(", left+1, right)
		}

		if right < left {
			backtrack(curString+")", left, right+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
