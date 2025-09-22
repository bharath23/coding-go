package leetcode0053

import "math"

/*
This is a simple two-pass alogrithm that evaluates all possible subarrays. The
time complexity is O(n^2), as we evaluate each subarray. Since no additional
data structures are used, the space complexity is O(1).

Time complexity: O(n^2)
Space complexity: O(1)
*/
func maxSubArrayv0(nums []int) int {
	maxSum := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

/*
We implement Kadane's alogrithm. At each index, decide if to start a
new subarray or extend the existing one. If the current subarray sum is
negative, discard it because it can not contribute to a maximum. Keep track of
the maximum sum found so far. Since this is one-pass algorithm the time
complexity is O(1). Since no additional data structures are used, the space
complexity is O(1).

Time complexity: O(n)
Space complexity: O(1)
*/
func maxSubArrayv1(nums []int) int {
	maxSum := math.MinInt
	currSum := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if currSum < 0 {
			currSum = n
		} else {
			currSum += n
		}

		if maxSum < currSum {
			maxSum = currSum
		}
	}

	return maxSum
}
