package main

//func main() {}

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	for i := range nums {
		sum += nums[i]
		// 避免全负数
		if sum > max {
			max = sum
		}
		// 换一个计数区间
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
