package leetcode0078

/*
 * The function generates all 2^n subsets, copying subsets takes O(n) time.
 * Hence, the time complexity is O(nx2^n)
 *
 * We store all the 2^n subsets with each of length up to n, requiring
 * additional space of O(nx2^n).
 *
 * Time complexity: O(n*2^n)
 *
 */
func subsetsv0(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 0 {
		return result
	}

	result = append(result, []int{})
	for _, n := range nums {
		subsets := [][]int{}
		for _, r := range result {
			subset := make([]int, len(r))
			copy(subset, r)
			subset = append(subset, n)
			subsets = append(subsets, subset)
		}

		for _, s := range subsets {
			result = append(result, s)
		}
	}

	return result
}

/*
 * The function generates all 2^n possbile subsets. Each subset requires O(n)
 * time to bild (copy elements). Overall time complexity is O(nx2^n).
 *
 * The recursion depth is O(n), the output stores all 2^n subsets, each up to
 * length of n, requiring O(nx2^n) space. So total space complexity is O(nx2^n).
 *
 * Time complexity: O(nx2^n)
 * Space complexity: O(nx2^n)
 */
func subsetsv1(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, cur []int)
	backtrack = func(start int, cur []int) {
		subset := make([]int, len(cur))
		copy(subset, cur)
		result = append(result, subset)
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

/*
 * The function generates all 2^n subsets, each subset requires O(n) time
 * build, checking n bits to decide which elements to include. Overall time
 * complexity is O(nx2^n).
 *
 * Output contains all subsets, totalling 2^n subsets with up to n elements
 * each. Space complexity is O(nx2^n).
 *
 * Time complexity: O(nx2^n)
 * Space complexity: O(nx2^n)
 */
func subsetsv2(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	result := make([][]int, 0, total)
	for mask := 0; mask < total; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}

		result = append(result, subset)
	}
	return result
}
