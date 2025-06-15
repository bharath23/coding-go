package leetcode0050

func myPowv0(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPowv0(x, -n)
	}

	res := 1.0
	for i := 0; i < n; i++ {
		res *= x
	}

	return res
}

func myPowv1(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1.0 / x
	}

	res := 1.0
	for n != 0 {
		if n%2 == 1 {
			res *= x
			n -= 1
		}

		x = x * x
		n = n / 2
	}

	return res
}
