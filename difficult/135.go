package main

//func main() {
//	fmt.Println(candy([]int{1, 3, 4, 5, 2}))
//}
func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// 每人至少一个
	for i := range candies {
		candies[i] = 1
	}

	// 左 -> 右
	// 确保右边孩子比左边相邻孩子的糖多
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 右 -> 左
	// 这里确保左边孩子比右边相邻孩子的糖多
	// 同样的，这里有一个贪心策略，就是需要拿最多的糖，确保大于左边相邻孩子和右边相邻孩子
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = Max(candies[i], candies[i+1]+1)
		}
	}

	sum := 0
	for i := range candies {
		sum += candies[i]
	}

	return sum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
