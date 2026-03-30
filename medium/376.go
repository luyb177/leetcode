package main

//func main() {}

// 需要考虑三种情况：
// 1. 上下坡度有平坡
// 2. 只有两个元素
// 3. 单调坡度出现平坡
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var prediff = 0
	var curdiff = 0
	var count = 1 // 默认序列最右有一个峰值

	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		// prediff = 0 说明之前是平坡，去除平坡左值
		if (prediff <= 0 && curdiff > 0) || (prediff >= 0 && curdiff < 0) {
			count++
			prediff = curdiff // 只在摆动的时候加
		}
	}
	return count
}
