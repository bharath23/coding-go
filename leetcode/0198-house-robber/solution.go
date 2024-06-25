package leetcode0198

func robV0(nums []int) int {
	memo := make(map[int]int)
	l := len(nums)
	if l > 0 {
		memo[0] = nums[0]
	} else {
		return 0
	}

	if l > 1 {
		memo[1] = max(nums[0], nums[1])
	} else {
		return memo[0]
	}

	var maxAmount func(int) int
	maxAmount = func(n int) int {
		if v, ok := memo[n]; ok {
			return v
		}

		memo[n] = max(maxAmount(n-1), maxAmount(n-2)+nums[n])
		return memo[n]
	}

	return maxAmount(len(nums) - 1)
}

func robV1(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[1], nums[0])
	}

	prev := nums[0]
	cur := max(nums[1], nums[0])
	for i := 2; i < l; i++ {
		cur, prev = max(prev+nums[i], cur), cur
	}

	return cur
}
