package main

//func main() {
//	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
//
//}

// 不需要排序
func minSubArrayLen1(target int, nums []int) int {
	left := 0
	func() {
		for i := 0; i < len(nums)-1; i++ {
			for j := 0; j < len(nums)-1-i; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}()
	temp := 0
	for ; left < len(nums); left++ {
		temp += nums[left]
		if temp >= target {
			left++
			return left
		}
	}
	return 0

}
func minSubArrayLen(target int, nums []int) int {

	i := 0 //指向起始位置 ，动态调整
	j := 0 //指向终止位置
	temp := 0
	result := len(nums) + 1

	for ; j <= len(nums)-1; j++ {
		temp += nums[j]
		for temp >= target { //这里有坑，就是使用if的话只执行一次，但是执行一次之后依然是大于的
			sub := j - i + 1
			if result >= sub {

				result = sub
			}
			temp -= nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		return 0

	} else {
		return result

	}
}
