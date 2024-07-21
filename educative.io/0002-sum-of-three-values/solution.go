package educative0002

import (
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		sum := target - nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			subSum := nums[l] + nums[r]
			if subSum == sum {
				return true
			}
			if subSum > sum {
				r--
			} else {
				l++
			}
		}
	}

	return false
}
