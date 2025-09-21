package leetcode0035

/*
 * The loop performs a standard binary search. In each iteration, the search
 * range [l..r] is halved by comparing the target with the middle element.
 * Since the range shrinks by a factor of two each time, the total number of
 * iterations is logarithmic in the input size. Overall time complexity is
 *  O(log n).
 *
 * The space complexity is O(1) because there is no additional space
 * requirement.
 *
 * Time complexity: O(log n)
 * Space complexity: O(1)
 */

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
