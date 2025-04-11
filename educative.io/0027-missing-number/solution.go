package educative0027

func findMissingNumber(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] < n && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i {
			return i
		}
	}

	return n
}
