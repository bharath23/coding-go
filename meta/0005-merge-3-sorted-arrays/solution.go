package meta0005

func merge(a, b, c []int) []int {
	i, j, k := 0, 0, 0
	result := make([]int, 0, len(a)+len(b)+len(c))
	for i < len(a) || j < len(b) || k < len(c) {
		min := int(^uint(0) >> 1)
		source := 0
		if i < len(a) && a[i] <= min {
			min = a[i]
			source = 1
		}
		if j < len(b) && b[j] <= min {
			min = b[j]
			source = 2
		}
		if k < len(c) && c[k] <= min {
			min = c[k]
			source = 3
		}

		result = append(result, min)
		switch source {
		case 1:
			i++
		case 2:
			j++
		case 3:
			k++
		}
	}

	return result
}
