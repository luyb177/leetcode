package main

func moveZeroes1(nums []int) {
	slow := 0
	fast := 0
	for ; fast <= len(nums)-1; fast++ {
		if nums[fast] != 0 {
			temp := nums[fast]
			nums[fast] = 0
			nums[slow] = temp
			slow++
		}
	}
}
func moveZeroes2(nums []int) {
	slow := 0
	fast := 0
	//先把非0挑出来
	for ; fast <= len(nums)-1; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	//最后一起赋值
	for ; slow <= len(nums)-1; slow++ {
		nums[slow] = 0
	}
}
