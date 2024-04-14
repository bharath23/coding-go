package leetcode0070

func climbStairsV0(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairsV0(n-1) + climbStairsV0(n-2)
}

func climbStairsV1(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	prev := 1
	cur := 2
	for i := 3; i <= n; i++ {
		cur, prev = cur+prev, cur
	}

	return cur
}

func climbStairsV2(n int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 2

	var noSteps func(int) int
	noSteps = func(n int) int {
		if v, ok := memo[n]; ok {
			return v
		}

		memo[n] = noSteps(n-1) + noSteps(n-2)
		return memo[n]
	}

	return noSteps(n)
}
