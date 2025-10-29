package main

//	func main() {
//		a := []int{-1, 0, 3, 5, 9, 12}
//		b := []int{-1, 0, 3, 5, 9, 12}
//		fmt.Println(search(a, 9))
//		fmt.Println(search(b, 2))
//	}
func search(nums []int, target int) int {
	//左闭右闭区间
	left := 0
	right := len(nums) - 1
	for left <= right { //此处<=是符合区间的
		mid := left + (right-left)/2
		if nums[mid] > target { //大于 更新右边界
			right = mid - 1 //已经不需要包含mid了
		} else if nums[mid] < target {
			left = mid + 1 //同上 不需要包含mid
		} else {
			return mid
		}
	}
	return -1
}
