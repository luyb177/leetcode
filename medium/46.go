package main

//func main() {}
//var path []int
//var res [][]int
//
//func permute(nums []int) [][]int {
//	path = []int{}
//	res = [][]int{}
//
//	used := make([]bool,len(nums))
//	backTracking(nums,used)
//	return res
//}
//
//func backTracking (nums []int,used []bool) {
//	if len(path) == len(nums) {
//		temp := make([]int,len(nums))
//		copy(temp,path)
//		res = append(res,temp)
//		return
//	}
//
//	for i:=0; i < len(nums); i++ {
//		if used[i] {
//			continue
//		}
//		path =append(path,nums[i])
//		used[i] = true
//		backTracking(nums,used)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
