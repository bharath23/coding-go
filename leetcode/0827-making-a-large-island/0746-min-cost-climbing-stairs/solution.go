package leetcode0746

func minCostClimbingStairsV0(cost []int) int {
	memo := make(map[int]int)
	var minCost func(int) int
	minCost = func(n int) int {
		if n < 0 {
			return 0
		}

		if n == 0 || n == 1 {
			memo[n] = cost[n]
			return cost[n]
		}

		if v, ok := memo[n]; ok {
			return v
		}

		memo[n] = min(minCost(n-1), minCost(n-2)) + cost[n]
		return memo[n]
	}

	l := len(cost)
	return min(minCost(l-1), minCost(l-2))
}

func minCostClimbingStairsV1(cost []int) int {
	l := len(cost)
	minCost := make(map[int]int)
	minCost[0] = cost[0]
	minCost[1] = cost[1]
	for i := 2; i < l; i++ {
		minCost[i] = min(minCost[i-1], minCost[i-2]) + cost[i]
	}

	return min(minCost[l-1], minCost[l-2])
}

func minCostClimbingStairsV2(cost []int) int {
	prev := cost[0]
	cur := cost[1]
	for i := 2; i < len(cost); i++ {
		prev, cur = cur, cost[i]+min(prev, cur)
	}

	return min(prev, cur)
}
