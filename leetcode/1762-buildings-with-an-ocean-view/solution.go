package leetcode1762

import "slices"

func findBuildings(heights []int) []int {
	result := []int{}
	maxHeight := -1
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			result = append(result, i)
			maxHeight = heights[i]
		}
	}

	slices.Reverse(result)
	return result
}
