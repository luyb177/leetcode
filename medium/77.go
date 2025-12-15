package main

//func main() {
//
//}
//
//// 组合
//var path []int
//var res [][]int
//
//func combine(n int, k int) [][]int {
//	res = make([][]int, 0)
//	path = make([]int, 0, k)
//
//	backtracking(n, k, 1)
//
//	return res
//}
//
//func backtracking(n int, k int, startIndex int) {
//	if len(path) == k {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//
//	for i := startIndex; i < n; i++ { // 修建一下的话 i < n - (k - len(path)) + 1
//		path = append(path, i)
//		backtracking(n, k, i+1)
//		path = path[:len(path)-1]
//	}
//}
