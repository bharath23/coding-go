package educative0028

func smallestMissingPositiveInteger(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}

	return n + 1
}
