package leetcode0121

func maxProfitv0(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			maxProfit = max(maxProfit, prices[j]-prices[i])
		}
	}

	return maxProfit
}

func maxProfitv1(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		}
	}
	return maxProfit
}
