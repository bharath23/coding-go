package meta0004

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func getMinCodeEntryTime(N, M int, C []int) int {
	cur := 1
	t := 0
	for _, c := range C {
		s := abs(c - cur)
		s = min(s, N-s)
		cur = c
		t += s
	}

	return t
}
