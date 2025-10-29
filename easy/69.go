package main

//func main() {
//	fmt.Println(mySqrt(8))
//}

// 很草率
func mySqrt1(x int) int {
	i := 1

	for ; i*i <= x; i++ {

	}
	return i - 1
}
func mySqrt(x int) int {
	left := 0
	right := x
	upper := 0
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			left = mid + 1
			upper = mid
		} else {
			right = mid - 1
		}
	}
	return upper
}
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
