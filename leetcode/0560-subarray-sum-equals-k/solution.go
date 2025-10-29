package leetcode0560

func subarraySumv0(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySumv1(nums []int, k int) int {
	sumCount := map[int]int{0: 1} // prefixSum -> count
	prefixSum := 0
	count := 0
	for _, num := range nums {
		prefixSum += num
		if freq, exists := sumCount[prefixSum-k]; exists {
			count += freq
		}
		sumCount[prefixSum]++
	}

	return count
}
