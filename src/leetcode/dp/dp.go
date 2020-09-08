package dp

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	prices[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > 0 {
			prices[i] = prices[i] - min
		} else {
			min = prices[i]
			prices[i] = 0
		}
	}
	max := prices[0]
	for _, v := range prices {
		if v > max {
			max = v
		}
	}
	return max
}
