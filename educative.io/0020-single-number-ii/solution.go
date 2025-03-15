package educative0020

func twoSingleNumbers(arr []int) []int {
	xorResult := 0
	for _, v := range arr {
		xorResult ^= v
	}

	mask := xorResult & -xorResult
	result := make([]int, 2)
	result[0] = 0
	for _, v := range arr {
		if v&mask != 0 {
			result[0] ^= v
		}
	}

	result[1] = result[0] ^ xorResult
	return result
}
