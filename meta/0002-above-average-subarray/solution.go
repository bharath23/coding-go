package meta0002

func aboveAverageSubarrays(array []int) [][2]int {
	result := [][2]int{}
	total := 0
	n := len(array)
	for _, a := range array {
		total += a
	}

	for i := 0; i < n; i++ {
		sub := 0
		for j := i; j < n; j++ {
			sub += array[j]
			if n == (j - i + 1) {
				result = append(result, [2]int{i + 1, j + 1})
			} else if sub/(j-i+1) > (total-sub)/(n-j+i-1) {
				result = append(result, [2]int{i + 1, j + 1})
			}
		}
	}

	return result
}
