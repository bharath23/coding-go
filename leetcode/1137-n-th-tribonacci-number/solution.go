package leetcode1137

func tribonacciV0(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 1
	var trib func(int) int
	trib = func(n int) int {
		if _, ok := memo[n]; !ok {
			memo[n] = trib(n-1) + trib(n-2) + trib(n-3)
		}

		return memo[n]
	}

	return trib(n)
}

func tribonacciV1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	trib := make([]int, n)
	trib[0] = 0
	trib[1] = 1
	trib[2] = 1
	for i := 3; i < n; i++ {
		trib[i] = trib[i-1] + trib[i-2] + trib[i-3]
	}

	return trib[n-1] + trib[n-2] + trib[n-3]
}

func tribonacciV2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	n0 := 0
	n1 := 1
	n2 := 1
	for i := 3; i < n; i++ {
		n0, n1, n2 = n1, n2, n0+n1+n2
	}

	return n0 + n1 + n2
}
