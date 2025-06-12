package leetcode0034

/*
We perform two binary searches: one to find the first occurrence of the target,
and another to find the last occurrence. Each binary search runs in O(log n)
time, so the total time complexity remains O(log n). Since we don’t use any
additional data structures, the space complexity is O(1).

Time complexity: O(logn)
Space complexity: O(1)
*/
func searchRange(nums []int, target int) []int {
	findBound := func(nums []int, target int, isStart bool) int {
		l := 0
		r := len(nums) - 1
		for l <= r {
			m := (l + r) / 2
			switch {
			case nums[m] < target:
				l = m + 1
			case nums[m] > target:
				r = m - 1
			case nums[m] == target:
				if isStart {
					if l == m || nums[m-1] != target {
						return m
					}
					r = m - 1
				} else {
					if r == m || nums[m+1] != target {
						return m
					}
					l = m + 1
				}
			}
		}

		return -1
	}

	start := findBound(nums, target, true)
	end := findBound(nums, target, false)
	return []int{start, end}
}
