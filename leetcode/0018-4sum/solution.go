package leetcode0018

import (
	"sort"
)

func twoSumv0(nums []int, start, target int) [][]int {
	res := [][]int{}
	lo := start
	hi := len(nums) - 1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if (sum < target) || (lo > start && nums[lo] == nums[lo-1]) {
			lo++
		} else if (sum > target) || ((hi < len(nums)-1) && nums[hi] == nums[hi+1]) {
			hi--
		} else {
			res = append(res, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}

	return res
}

/*
 * We recursively fix one element at each level and reduce the problem to a
 * smaller k-1 sum. When we reach k = 2, we solve it using the two-pointer
 * approach in linear time. At each recursive level, we loop through the array
 * starting from a given index. The worst-case time complexity is O(n^(k-1)).
 * Since we use recursion, the maximum depth is k, so the space is O(k).
 *
 * Time complexity: O(n^(k-1))
 * Space complexity: O(k)
 */
func kSumv0(nums []int, target, start, k int) [][]int {
	res := [][]int{}
	if start == len(nums) {
		return res
	}

	avg := target / k
	if nums[start] > avg || nums[len(nums)-1] < avg {
		return res
	}

	if k == 2 {
		return twoSumv0(nums, start, target)
	}

	for i := start; i < len(nums)-i; i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSumv0(nums, target-nums[i], i+1, k-1) {
				res = append(res, append([]int{nums[i]}, subset...))
			}
		}
	}

	return res
}

func fourSumv0(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSumv0(nums, target, 0, 4)
}

func twoSumv1(nums []int, start, target int) [][]int {
	res := [][]int{}
	complements := make(map[int]bool, len(nums))
	for i := start; i < len(nums); i++ {
		if len(res) == 0 || res[len(res)-1][1] != nums[i] {
			if complements[target-nums[i]] {
				res = append(res, []int{target - nums[i], nums[i]})
			}
		}
		complements[nums[i]] = true
	}
	return res
}

/*
* We recursively fix one element at each level and reduce the problem to a
* smaller k-1 sum. When we reach k = 2, we use a hash map to find pairs that
* sum to the target in linear time. At each recursive level, we loop through
* the array starting from a given index. The worst-case time complexity is
* O(n^(k-1)).
* Since we use recursion, the maximum depth is k, so the space is O(k).At the
* two sum level we use a hash map which requires O(n) space.
*
* Time complexity: O(n^(k-1))
* Space complexity: O(n+k)
 */
func kSumv1(nums []int, target, start, k int) [][]int {
	res := [][]int{}
	if start == len(nums) {
		return res
	}

	avg := target / k
	if nums[start] > avg || nums[len(nums)-1] < avg {
		return res
	}

	if k == 2 {
		return twoSumv1(nums, start, target)
	}

	for i := start; i < len(nums)-i; i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSumv1(nums, target-nums[i], i+1, k-1) {
				res = append(res, append([]int{nums[i]}, subset...))
			}
		}
	}

	return res
}

func fourSumv1(nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSumv1(nums, target, 0, 4)
}

/*
* We use two nested loops to fix the first two elements, and then apply the
* two-pointer technique to find the remaining two elements. This results in a
* time complexity of O(n^3) in the worst case, where n is the length of the
* input array.
* As for space complexity, apart from the output list, we do not use any
* additional data structures. Therefore, space complexity is O(1).
*
* Time complexity: O(n^3)
* Space complexity: O(k)
 */
func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l := j + 1
			r := n - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					result = append(result,
						[]int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return result
}
