package leetcode0031

func nextPermutation(nums []int) {
	pivot := len(nums) - 1
	for pivot > 0 && nums[pivot] <= nums[pivot-1] {
		pivot--
	}

	if pivot != 0 {
		i := len(nums) - 1
		for nums[i] <= nums[pivot-1] {
			i--
		}

		nums[pivot-1], nums[i] = nums[i], nums[pivot-1]
	}

	l := pivot
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
