package main

//func main() {}
//var path []int
//var res [][]int
//
//func subsets(nums []int) [][]int {
//	path = []int{}
//	res = [][]int{}
//
//	backTracking(nums,0)
//	return res
//}
//
//func backTracking (nums []int, startIndex int) {
//	if startIndex <= len(nums) {
//		temp := make([]int,len(path))
//		copy(temp,path)
//		res = append(res,temp)
//	}
//
//	for i := startIndex; i < len(nums);i++ {
//		path = append(path,nums[i])
//		backTracking(nums,i+1)
//		path = path[:len(path)-1]
//	}
//}
