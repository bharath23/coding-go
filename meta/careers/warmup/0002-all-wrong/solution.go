package meta0002

func getWrongAnswers(N int, C string) string {
	runeA := 'A'
	runeB := 'B'
	result := make([]rune, N)
	for i, c := range C {
		if c == runeA {
			result[i] = runeB
		} else {
			result[i] = runeA
		}
	}

	return string(result)
}
