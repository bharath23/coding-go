package educative0001

func isPalindrome(inputString string) bool {
	l := 0
	r := len(inputString) - 1
	for l <= r {
		if inputString[l] != inputString[r] {
			return false
		}

		l++
		r--
	}

	return true
}
