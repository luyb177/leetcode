package main

//	func main() {
//		fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
//		fmt.Println(spiralOrder([][]int{{7}, {9}, {5}}))
//		fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24}}))
//	}
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	nums := make([]int, m*n)

	p := min(m, n)
	x := 0
	y := 0
	sub := 1
	count := 0
	for temp := 0; temp < p/2; temp++ {
		i := x
		j := y
		for ; j < n-sub; j++ { // 上行
			nums[count] = matrix[i][j]
			count++
		}
		for ; i < m-sub; i++ { //右列
			nums[count] = matrix[i][j]
			count++
		}
		for ; j > y; j-- {
			nums[count] = matrix[i][j]
			count++
		}
		for ; i > x; i-- {
			nums[count] = matrix[i][j]
			count++
		}
		x++
		y++
		sub++
	}
	if m >= n {
		if n%2 == 1 {
			for ; x <= m-sub; x++ {
				nums[count] = matrix[x][y]
				count++
			}
		}
	} else {
		if m%2 == 1 {
			for ; y <= n-sub; y++ {
				nums[count] = matrix[x][y]
				count++
			}
		}
	}

	return nums
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
