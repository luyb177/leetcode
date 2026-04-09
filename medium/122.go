package main

//func main() {}

func maxProfit(prices []int) int {
	// 分解利润
	// price[3] - price[0] == price[3] - price[2] + price[2] - price[1] + price[1] - price[0]
	// 今日买 明日卖，计算利润
	total := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			total += prices[i+1] - prices[i]
		}
	}
	return total
}
