package greedy

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var res int
	for i := 1; i < len(prices); i++ {
		sub := prices[i] - prices[i-1]
		if sub > 0 {
			res += sub
		}
	}
	return res
}
