package main

// func main() {
//
//		result := generateMatrix(4)
//		for i := range result {
//			for j := range result[i] {
//				fmt.Print(result[i][j])
//			}
//			fmt.Println()
//		}
//	}
func generateMatrix(n int) [][]int { //左闭右开
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n) //分配内存空间
	}

	x := 0
	y := 0
	sub := 1
	count := 1

	for temp := 0; temp < n/2; temp++ { //循环n/2圈
		i := x
		j := y
		for ; j < n-sub; j++ {
			nums[i][j] = count //上行
			count++
		}
		for ; i < n-sub; i++ {
			nums[i][j] = count //右列
			count++
		}
		for ; j > y; j-- {
			nums[i][j] = count //下行
			count++
		}
		for ; i > x; i-- {
			nums[i][j] = count
			count++ //左行
		}
		//循环一圈后
		x++
		y++
		sub++
	}
	if n%2 == 1 {
		nums[x][y] = count
	}
	return nums
}
