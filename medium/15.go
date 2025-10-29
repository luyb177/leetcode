package main

//func main() {
//	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
//}

//暴力解法暂时没过去

//func threeSum(nums []int) [][]int {
//	result := make([][]int, 0)
//	if len(nums) < 3 {
//		return result
//	}
//	//先排序
//	for i := 0; i < len(nums)-1; i++ {
//		for j := 0; j < len(nums)-1-i; j++ {
//			if nums[i] < nums[j] {
//				nums[i], nums[j] = nums[j], nums[i]
//			}
//		}
//	}
//	for left := 0; left < len(nums)-2; left++ {
//		for mid := left + 1; mid < len(nums)-1; mid++ {
//			for right := mid + 1; right < len(nums); right++ {
//				if nums[left]+nums[mid]+nums[right] == 0 {
//					result = append(result, []int{nums[left], nums[mid], nums[right]})
//				}
//			}
//		}
//	}
//	return result
//}

func threeSum(nums []int) [][]int {

	// 排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	var result [][]int

	// 双指针
	for i := 0; i < len(nums)-2; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 判断
		if nums[i] > 0 { // 已经大于0了，没必要继续循环
			break
		}

		var left = i + 1
		var right = len(nums) - 1

		for left < right { // 不包含 =
			sum := nums[i] + nums[left] + nums[right]
			// 找到三个数
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}

	}
	return result
}
