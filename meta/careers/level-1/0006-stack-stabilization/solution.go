package meta0006

func getMinimumDeflatedDiscCount(N int, R []int) int {
	num := 0
	for i := N - 1; i >= 0; i-- {
		if R[i] < i+1 {
			return -1
		}

		if i < N-1 && R[i] >= R[i+1] {
			R[i] = R[i+1] - 1
			num++
		}
	}

	return num
}
