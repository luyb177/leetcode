package main

//	func main() {
//		fmt.Println(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
//	}
func fourSum(nums []int, target int) [][]int {
	// 对数组进行排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// 三指针 - > 双指针
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		// 提前判断
		if nums[i] > target && (nums[i] >= 0 || target >= 0) {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left := i + 1; left < len(nums)-2; left++ {
			// 提前判断
			if nums[i]+nums[left] > target && nums[i]+nums[left] >= 0 {
				break
			}

			// 去重
			if left > i+1 && nums[left] == nums[left-1] { // 相当于一个新的数组，left 要大于0的时候才可以防止越界。这里的0对应的下标是i+1
				continue
			}

			// 双指针
			var mid = left + 1
			var right = len(nums) - 1
			for mid < right {
				sum := nums[i] + nums[left] + nums[mid] + nums[right]

				if sum == target {
					// 添加结果
					result = append(result, []int{nums[i], nums[left], nums[mid], nums[right]})

					// 去重
					for mid < right && nums[mid] == nums[mid+1] {
						mid++
					}
					for mid < right && nums[right] == nums[right-1] {
						right--
					}
					mid++
					right--
				} else if sum > target {
					right--
				} else {
					mid++
				}
			}
		}

	}
	return result
}
