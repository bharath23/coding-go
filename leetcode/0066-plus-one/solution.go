package leetcode0066

/*
 * The algorithm processes the digits from right to left. If the current digit
 * is less than 9, it increments that digit and returns immediately. If the
 * digit is 9, it is set to 0 and the carry continues to the next digit. In the
 * case where all digits are 9, the loop completes and a new leading 1 is
 * added.
 *
 * In the worst case, every digit is processed once, giving O(n) time
 * complexity for an input array of length n. Appending a new leading digit
 * also takes O(n) time since it creates a new slice.
 *
 * The space complexity is O(n) in the all-9s case due to allocating a new
 * slice of size n+1, otherwise O(1) extra space.
 *
 * Time complexity: O(n)
 * Space complexity: O(n) in the worst case, O(1) otherwise
 */
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	digits = append([]int{1}, digits...)
	return digits
}
