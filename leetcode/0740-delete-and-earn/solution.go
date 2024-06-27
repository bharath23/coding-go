package leetcode0740

func deleteAndEarnV0(nums []int) int {
	points := make(map[int]int)
	maxNum := 0
	for _, n := range nums {
		points[n] += n
		maxNum = max(maxNum, n)
	}

	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = points[1]
	var maxPoints func(int) int
	maxPoints = func(n int) int {
		if _, ok := memo[n]; !ok {
			memo[n] = max(maxPoints(n-1), maxPoints(n-2)+points[n])
		}

		return memo[n]
	}

	return maxPoints(maxNum)
}

func deleteAndEarnV1(nums []int) int {
	points := make(map[int]int)
	maxNum := 0
	for _, n := range nums {
		points[n] += n
		maxNum = max(maxNum, n)
	}

	maxPoints := make([]int, maxNum+1)
	maxPoints[1] = points[1]
	for i := 2; i <= maxNum; i++ {
		maxPoints[i] = max(maxPoints[i-1], maxPoints[i-2]+points[i])
	}

	return maxPoints[maxNum]
}

func deleteAndEarnV2(nums []int) int {
	points := make(map[int]int)
	maxNum := 0
	for _, n := range nums {
		points[n] += n
		maxNum = max(maxNum, n)
	}

	prev := 0
	cur := points[1]
	for i := 2; i <= maxNum; i++ {
		prev, cur = cur, max(cur, prev+points[i])
	}

	return cur
}
