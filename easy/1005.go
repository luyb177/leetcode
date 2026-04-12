package main

//func main()
func largestSumAfterKNegations(nums []int, k int) int {
	// 按绝对值从大到小排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if abs(nums[j]) < abs(nums[j+1]) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// 遍历，将负数给反转
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}

	// k 如果不等于0 将最小数反复取反
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}

	// 求和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
