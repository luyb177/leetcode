package main

//	func main() {
//		a := []int{1, 3, 5, 6}
//		fmt.Println(searchInsert(a, 5))
//		fmt.Println(searchInsert(a, 2))
//		fmt.Println(searchInsert(a, 7))
//
// }
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1 //右移一位
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid //找到的下标
		}
	}

	//没找到但退出循环了 证明 left > right
	return left
}
