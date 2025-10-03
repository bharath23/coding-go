package leetcode1249

/*
 * The solution uses three loops:
 * 1. The first loop traverses the entire string.
 * 2. The second loop processes the stack and records indices in a map.
 * 3. The third loop builds the final result string.
 *
 * Each loop runs in O(n), so the overall time complexity is O(n).
 *
 * Additional data structures such as a stack and a map are used,
 * each requiring up to O(n) space. Thus, the overall space complexity is O(n).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(n)
 */
func minRemoveToMakeValidv0(s string) string {
	result := ""
	stack := []int{}
	parenthesesToRemove := make(map[int]bool)
	for i, c := range s {
		if c != '(' && c != ')' {
			continue
		}

		if c == '(' {
			stack = append(stack, i)
		} else if len(stack) == 0 {
			parenthesesToRemove[i] = true
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	for _, idx := range stack {
		parenthesesToRemove[idx] = true
	}

	for i, c := range s {
		if parenthesesToRemove[i] {
			continue
		}

		result += string(c)
	}

	return result
}

/*
 * The solution makes three passes over the string:
 * 1. The first pass (left to right) collects valid ')' and all other
 * characters.
 * 2. The second pass (right to left) removes excess '('.
 * 3. The third pass reverses the result to restore the correct order.
 *
 * Each pass processes at most n characters, so the overall time complexity is
 * O(n).
 *
 * Additional space is used for the intermediate slices (firstPass, result),
 * each storing up to n characters. Thus, the space complexity is O(n).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(n)
 */
func minRemoveToMakeValidv1(s string) string {
	firstPass := []rune{}
	balanced := 0
	for _, c := range s {
		switch c {
		case '(':
			balanced++
			firstPass = append(firstPass, c)
		case ')':
			if balanced > 0 {
				balanced--
				firstPass = append(firstPass, c)
			}
		default:
			firstPass = append(firstPass, c)
		}
	}

	result := []rune{}
	balanced = 0
	for i := len(firstPass) - 1; i >= 0; i-- {
		switch r := firstPass[i]; r {
		case ')':
			balanced++
			result = append(result, r)
		case '(':
			if balanced > 0 {
				balanced--
				result = append(result, r)
			}
		default:
			result = append(result, r)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

/*
 * The solution makes three passes over the string:
 * 1. The first pass (left to right) collects valid ')' and all other
 *  characters, while tracking the number of unmatched '('.
 * 2. The second pass (right to left) removes any remaining extra '('.
 * 3. The third pass reverses the result to restore the original order.
 *
 * Each pass processes at most n characters, so the overall time complexity is
 * O(n).
 *
 * Additional space is used for the intermediate slices (firstPass and result),
 * each storing up to n characters. Thus, the space complexity is O(n).
 *
 * Time Complexity:  O(n)
 * Space Complexity: O(n)
 */
func minRemoveToMakeValidv2(s string) string {
	firstPass := []rune{}
	balanced := 0
	for _, c := range s {
		switch c {
		case '(':
			balanced++
			firstPass = append(firstPass, c)
		case ')':
			if balanced > 0 {
				balanced--
				firstPass = append(firstPass, c)
			}
		default:
			firstPass = append(firstPass, c)
		}
	}

	result := []rune{}
	for i := len(firstPass) - 1; i >= 0; i-- {
		if c := firstPass[i]; c == '(' {
			if balanced > 0 {
				balanced--
			} else {
				result = append(result, c)
			}
		} else {
			result = append(result, c)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
