package main

// func main() {}
func findContentChildren(g []int, s []int) int {
	// 小饼干优先喂小胃口的人
	bubbleSort(g)
	bubbleSort(s)
	var count = 0

	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			// 饼干大就给
			count++
			i++
			j++
		} else {
			// 饼干小就下一个饼干
			j++
		}
	}
	return count
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		flag := false

		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		if !flag {
			return
		}
	}
}
