package leetcode0219

/*
This is a straightforward two-pass sliding window implementation. The time
complexity is O(n × min(k, n)), which simplifies to O(nk) when k < n, and up to
O(n²) in the worst case. The space complexity is O(1), as no additional data
structures are used.

Time complexity: O(n^2)
Space complexity: O(1)
*/
func containsNearbyDuplicatev0(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := max(i-k, 0); j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

/*
This is a single-pass sliding window implementation with a hash set for
constant-time lookups. The time complexity is O(n), as each element is added
and removed from the set at most once. The space complexity is O(min(k, n)),
since the set maintains at most k elements at any given time. In the worst
case, this simplifies to O(n).

Time complexity: O(n)
Space complexity: O(n)
*/
func containsNearbyDuplicatev1(nums []int, k int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		isSeen := seen[nums[i]]
		if isSeen {
			return true
		}

		seen[nums[i]] = true
		if i >= k {
			delete(seen, nums[i-k])
		}
	}

	return false
}
