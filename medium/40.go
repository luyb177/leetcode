package main

// func main() {}
//var path []int
//var res [][]int
//
//// 因为时间原因，所以需要边查询边去重
//
//func combinationSum2(candidates []int, target int) [][]int {
//	path = []int{}
//	res = [][]int{}
//
//	// 排序，方便后续递归去重
//	for i := 0; i < len(candidates)-1; i++ {
//		swapped := false
//		for j := 0; j < len(candidates)-i-1; j++ {
//			if candidates[j] < candidates[j+1] {
//				candidates[j], candidates[j+1] = candidates[j+1], candidates[j]
//				swapped = true
//			}
//		}
//		if !swapped {
//			break
//		}
//	}
//
//	used := make([]bool, len(candidates))
//
//	backTracking(candidates, target, 0, used)
//	return res
//}
//
//func backTracking(candidates []int, target int, startIndex int, used []bool) {
//	if target < 0 {
//		return
//	}
//
//	if target == 0 {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//	}
//
//	for i := startIndex; i < len(candidates); i++ {
//		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
//			// 说明是同一树层在用
//			continue
//		}
//		path = append(path, candidates[i])
//		used[i] = true
//		backTracking(candidates, target-candidates[i], i+1, used)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
