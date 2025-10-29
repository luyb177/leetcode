package main

//func main() {
//	fmt.Println(threeSumClosest([]int{-1000, -1000, -1000}, 10000))
//}

// 非常暴力
// 特例
func threeSumClosest(nums []int, target int) int {
	result := 10000
	temp := 10000
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	for left := 0; left < len(nums)-2; left++ {
		for mid := left + 1; mid < len(nums)-1; mid++ {
			for right := mid + 1; right < len(nums); right++ {

				temp1 := nums[left] + nums[mid] + nums[right] - target
				if abs(temp) >= abs(temp1) {
					result = nums[left] + nums[mid] + nums[right]
					temp = temp1
				}

			}
		}
	}
	return result
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
