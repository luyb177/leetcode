package main

//func main() {}

//var path []int
//var result [][]int
//
//func subsetsWithDup(nums []int) [][]int {
//	path = []int{}
//	result = [][]int{}
//
//	// 将nums排序
//	for i := 0; i < len(nums) - 1; i ++ {
//		flag := false
//		for j := 0; j < len(nums) - i - 1;j++ {
//			if nums[j] > nums[j+1] {
//				flag = true
//				nums[j], nums[j+1] = nums[j+1], nums[j]
//			}
//		}
//		if !flag  {
//			break
//		}
//	}
//
//	used := make([]bool,len(nums))
//	backTracking(nums,0,used)
//	return result
//}
//
//func backTracking(nums []int, startIndex int, used []bool) {
//	if startIndex <= len(nums) {
//		temp := make([]int,len(path))
//		copy(temp,path)
//		result = append(result,temp)
//	}
//
//	for i := startIndex; i < len(nums); i++ {
//		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
//			// 树层去重
//			continue
//		}
//		path = append(path,nums[i])
//		used[i] = true
//		backTracking(nums,i+1,used)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
