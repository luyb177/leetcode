package main

//	func main() {
//		fmt.Println(sortedSquares([]int{-5, -3, -2, -1}))
//	}
func sortedSquares1(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		k = i
		for j := i + 1; j < len(nums); j++ {
			if nums[k] > nums[j] {
				k = j
			}
		}
		temp := nums[i]
		nums[i] = nums[k]
		nums[k] = temp
	}
	return nums
}
func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	k := right
	//创建一个新数组
	result := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			result[k] = nums[right] * nums[right]
			right--
			k--
		} else {
			result[k] = nums[left] * nums[left]
			left++
			k--
		}
	}
	return result
}
