package stack

func MaxProfit(prices []int) int {
	var max, min, result int
	for i, price := range prices {
		if i == 0 {
			max, min = price, price
		}
		if price > max {
			result = maximum(result, price-min)
		}
		if price < min {
			max, min = price, price
		}
	}
	return result
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
