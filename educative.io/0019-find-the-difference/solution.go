package educative0019

func extraCharacterIndex(str1 string, str2 string) int {
	result := 0
	len1 := len(str1)
	len2 := len(str2)
	for i := 0; i < max(len1, len2); i++ {
		if i < len1 {
			result ^= int(str1[i])
		}

		if i < len2 {
			result ^= int(str2[i])
		}
	}

	var str string
	if len1 > len2 {
		str = str1
	} else {
		str = str2
	}

	for i := 0; i < max(len1, len2); i++ {
		if result == int(str[i]) {
			return i
		}
	}

	return -1
}
