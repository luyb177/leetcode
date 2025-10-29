package main

//func main() {
//	a := []int{0, 0, 0, 0, 0, 0, 0}
//	fmt.Println(Lowwer(a, 0), upper(a, 0))
//}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else { //找到一个等于的了
			//找左边的值
			leftTerm := mid - 1
			for ; leftTerm >= left; leftTerm-- {
				if nums[leftTerm] == target {
					continue
				} else {
					break
				}
			}
			rightTerm := mid + 1
			for ; rightTerm <= right; rightTerm++ {
				if nums[rightTerm] == target {
					continue
				} else {
					break
				}
			}
			return []int{leftTerm + 1, rightTerm - 1}

		}
	}
	return []int{-1, -1}
}

// 也可以二分查找左右边界
// 寻求左边界 >= 找到
func Lowwer(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	lowwer := -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target { //此时target 在左边界
			right = mid - 1
			lowwer = mid //更新一下
		} else { //此时 nums[mid] < target
			left = mid + 1
		}
	}
	return lowwer
}

// 寻找右边界 <=找到
func upper(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	upper := -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else { //此时 <= target
			left = mid + 1
			upper = mid
		}
	}
	return upper
}
