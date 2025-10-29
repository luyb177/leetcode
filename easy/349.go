package main

//func main() {
//
//}
func intersection(nums1 []int, nums2 []int) []int {
	//使用map 模拟set
	set := make(map[int]struct{},0)
	result := make([]int,0)
	for _,v:=range nums1 {
		if _,ok := set[v]; !ok{//记录nums1中存在的值
			set[v] = struct{}{} //使用空结构体 节省空间
		}
	}
	for _,v := range nums2 {
		if _,ok := set[v]; ok {
			//有交集
			result = append(result,v)
			//手动去重
			delete(set,v)
		}
	}
	return result
}
