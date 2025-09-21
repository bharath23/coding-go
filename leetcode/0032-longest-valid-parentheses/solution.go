package leetcode0032

/*
 * The loop iterates once through the string. Each dp[i] is computed in constant
 * time using previous values. Thus the overall time complexity is O(n).
 *
 * The space complexity is O(n) due to the dp array storing the longest valid
 * substring length ending at each index.
 *
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func longestValidParenthesesv0(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}

		if s[i-1] == '(' {
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else {
			idx := i - dp[i-1] - 1
			if idx >= 0 && s[idx] == '(' {
				dp[i] = dp[i-1] + 2
				if idx-1 >= 0 {
					dp[i] += dp[idx-1]
				}
			}
		}

		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

/*
 * The loop iterates through the string once. Each index is pushed to or popped
 * from the stack at most once, so the total stack operations are linear.
 * Lengths of valid substrings are calculated in constant time after each pop.
 * Overall time complexity is O(n).
 *
 * The space complexity is O(n) in the worst case, when all characters are '('
 * and thus stored in the stack.
 *
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func longestValidParenthesesv1(s string) int {
	result := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			length := i - stack[len(stack)-1]
			result = max(result, length)
		}
	}

	return result
}

/*
 * The string is scanned twice, once from left to right and once from right to
 * left. Each scan is linear, and constant-time updates are performed for each
 * character. Overall time complexity is O(n).
 *
 * The space complexity is O(1) since only two counters (left, right) and one
 * result variable are maintained.
 *
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func longestValidParenthesesv2(s string) int {
	result := 0

	// Scan left to right
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if result < 2*right {
				result = 2 * right
			}
		}
		if right > left {
			left, right = 0, 0
		}
	}

	// Scan from right to left
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			if result < 2*right {
				result = 2 * right
			}
		}
		if left > right {
			left, right = 0, 0
		}
	}

	return result
}
