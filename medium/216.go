package main

//func main() {
//
//}
//
//// 组合总和
//var res [][]int
//var path []int
//
//func combinationSum3(k int, n int) [][]int {
//	path = make([]int, 0, k)
//	res = make([][]int, 0)
//
//	backtracking(k, n, 1, 0)
//	return res
//}
//
//func backtracking(k int, n int, startIndex int, sum int) {
//	if sum == n && len(path) == k {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//	}
//
//	for i := startIndex; i <= 9; i++ {
//		path = append(path, i)
//		backtracking(k, n, i+1, sum+i)
//		path = path[:len(path)-1]
//	}
//}
