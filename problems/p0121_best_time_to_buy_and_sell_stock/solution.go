package p0121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	l := 0
	r := 1

	maxPrefit := max(0, prices[r]-prices[l])

	for ; r < len(prices); r++ {

		if prices[r] > prices[l] {
			maxPrefit = max(maxPrefit, prices[r]-prices[l])
		} else {
			l = r
		}
	}
	return maxPrefit
}
