package meta0001

func answerQueries(queries [][2]int, n int) []int {
	boolArray := make([]bool, n)
	result := []int{}
	for _, query := range queries {
		t, i := query[0], query[1]
		if t == 1 {
			boolArray[i-1] = true
		}
		if t == 2 {
			val := -1
			for j := i - 1; j < n; j++ {
				if boolArray[j] == true {
					val = j + 1
					break
				}
			}
			result = append(result, val)
		}
	}

	return result
}
