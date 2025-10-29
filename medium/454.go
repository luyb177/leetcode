package main

// func main() {
//
// }
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int) // key-a+b  value-a+b的次数
	result := 0
	// 从前两个数组中找 a+b
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if _, ok := m[v1+v2]; ok {
				m[v1+v2]++
			} else {
				m[v1+v2] = 1
			}
		}
	}

	// 后两个数组中找 v1+v2+v3+v4 = 0 即找 -v3-v4
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if _, ok := m[-(v3 + v4)]; ok {
				result += m[-(v3 + v4)]
			}
		}
	}
	return result
}
