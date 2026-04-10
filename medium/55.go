package main

// func main() {}
func canJump(nums []int) bool {
	// 局部贪心，看最大可跳跃范围
	cover := 0
	for i := 0; i <= cover; i++ {
		// 当前位置加上当前位置跳跃范围
		// 如果超过了之前的最大可跳跃范围，就更新最大可跳跃范围
		// 如果最大可跳跃范围超过了最后一个位置，说明可以跳到最后一个位置，返回true
		// 如果当前索引超过了最大可跳跃范围，说明无法跳到最后一个位置，返回false
		temp := i + nums[i]
		if temp > cover {
			cover = temp
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
