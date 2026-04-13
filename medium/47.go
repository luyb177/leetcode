package main

//func main() {}

//var path []int
//var res [][]int
//
//func permuteUnique(nums []int) [][]int {
//	path = []int{}
//	res = [][]int{}
//
//	// 排序为了方便去重
//	for i := 0; i < len(nums)-1; i++ {
//		flag := false
//		for j := 0; j < len(nums)-i-1; j++ {
//			if nums[j] > nums[j+1] {
//				flag = true
//				nums[j], nums[j+1] = nums[j+1], nums[j]
//			}
//		}
//
//		if !flag {
//			break
//		}
//	}
//	used := make([]bool, len(nums))
//	backTracking(nums, used)
//
//	return res
//}
//
//func backTracking(nums []int, used []bool) {
//	if len(path) == len(nums) {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		// 树层去重
//		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
//			continue
//		}
//
//		// 树枝去重
//		if used[i] {
//			continue
//		}
//		path = append(path, nums[i])
//		used[i] = true
//		backTracking(nums, used)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
