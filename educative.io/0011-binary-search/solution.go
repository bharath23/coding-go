package educative0011

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		switch {
		case nums[m] == target:
			return m
		case nums[m] < target:
			l = m + 1
		default:
			r = m - 1
		}
	}
	return -1
}
