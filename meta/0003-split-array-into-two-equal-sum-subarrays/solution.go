package meta0003

func canSplitArray(nums []int) int {
	leftSum := 0
	for _, v := range nums {
		leftSum += v
	}

	rightSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum += nums[i]
		leftSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
