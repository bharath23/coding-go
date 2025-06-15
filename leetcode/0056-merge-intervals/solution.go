package leetcode0056

import (
	"sort"
)

/*
We perform a single pass over the list in addition to sorting it. Therefore,
the overall time complexity is O(n log n), dominated by the sort operation. If
we use an in-place sort, the space complexity is O(1); otherwise, it depends on
the space requirements of the sorting algorithm used.

Time complexity: O(nlogn)
Space complexity: O(1)
*/
func merge(intervals [][]int) [][]int {
	merged := [][]int{}
	if len(intervals) == 0 {
		return merged
	}

	// Sort the intervals with start value of each interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	interval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if (interval[0] <= intervals[i][0] && intervals[i][0] <= interval[1]) ||
			intervals[i][0] <= interval[1] && interval[1] <= intervals[i][1] {
			interval[0] = min(interval[0], intervals[i][0])
			interval[1] = max(interval[1], intervals[i][1])
		} else {
			merged = append(merged, interval)
			interval = intervals[i]
		}
	}

	merged = append(merged, interval)

	return merged
}
