package meta0005

func getMinProblemCount(N int, S []int) int {
	problems := map[string]int{
		"1": 0,
		"2": 0,
	}

	for _, s := range S {
		twos := s / 2
		ones := s % 2
		if problems["2"] < twos {
			problems["2"] = twos
		}
		if ones == 1 {
			problems["1"] = ones
		}
	}

	return problems["1"] + problems["2"]
}
