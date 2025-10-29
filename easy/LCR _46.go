package main

// func main() {
//
// }
func spiralArray(array [][]int) []int {

	m := len(array)
	if m == 0 {
		return []int{}
	}
	n := len(array[0])
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
			nums[count] = array[i][j]
			count++
		}
		for ; i < m-sub; i++ { //右列
			nums[count] = array[i][j]
			count++
		}
		for ; j > y; j-- {
			nums[count] = array[i][j]
			count++
		}
		for ; i > x; i-- {
			nums[count] = array[i][j]
			count++
		}
		x++
		y++
		sub++
	}
	if m >= n {
		if n%2 == 1 {
			for ; x <= m-sub; x++ {
				nums[count] = array[x][y]
				count++
			}
		}
	} else {
		if m%2 == 1 {
			for ; y <= n-sub; y++ {
				nums[count] = array[x][y]
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
