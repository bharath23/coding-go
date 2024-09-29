package educative0009

func permutatePalindrome(st string) bool {
	letterFreq := make(map[int]int, 26)
	for _, c := range st {
		letterFreq[int(c-'a')] = letterFreq[int(c-'a')] + 1
	}

	numOdd :=0
	for _, freq := range letterFreq {
		if freq % 2 != 0{
			numOdd++
		}
	}

	if numOdd > 1 {
		return false
	}

	return true
}
