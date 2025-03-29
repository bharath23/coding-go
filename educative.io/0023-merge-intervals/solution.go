package educative0023

func mergeIntervals(intervals [][]int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		return result
	}

	result = append(result, intervals[0])
	cur := 0
	for _, interval := range intervals[1:] {
		if (result[cur][0] <= interval[0]) && (interval[0] <= result[cur][1]) {
			result[cur][0] = min(result[cur][0], interval[0])
			result[cur][1] = max(result[cur][1], interval[1])
		} else {
			result = append(result, interval)
			cur++
		}
	}

	return result
}
