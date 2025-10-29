package main

//func main() {
//	nums := []int{2, 7, 11, 15}
//	fmt.Println(twoSum(nums, 9))
//}

// 暴力解法
//
//	func twoSum(nums []int, target int) []int {
//		i := 0
//		j := 0
//		flag := false
//		for i = 0; i < len(nums); i++ {
//			p := nums[i]
//			for j = i + 1; j < len(nums); j++ {
//				q := nums[j]
//				if p+q == target {
//					flag = true
//					break
//				}
//			}
//			if flag {
//				break
//			}
//		}
//		return []int{i, j}
//	}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		x := target - num
		if index, ok := m[x]; ok {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}
