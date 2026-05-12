package main

//func main() {
//	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
//}

// Note: 这种获取每一个字符的最远出现索引，并判断当前字符串所拥有的最远出现索引
// 更快捷
func partitionLabels(s string) []int {
	m := make(map[int32]int)
	res := make([]int, 0)
	// 记录每一个字符的最远位置
	for i, v := range s {
		m[v] = i
	}

	// 找到当前字符串的最远边界
	left := 0
	right := 0
	for i, v := range s {
		index := m[v]
		if index > right {
			right = index
		}
		if i == right {
			res = append(res, right-left+1)
			right = i + 1
			left = i + 1
		}
	}
	return res
}

// 动态合并区间的贪心思想
// 但是时间复杂度比较高
// 使用哈希表维护字符所属分段，
// 遇到重复字符时动态合并分段，
// 本质属于区间合并的贪心策略。
// 更简单的是记录每个字符最后出现的位置
//func partitionLabels(s string) []int {
//	res := make([]int, 0)
//	// 记录 这个字母上次出现的位置
//	m := make(map[int32]int)
//	// 记录 这个位置的 字母
//	letter := make([][]int32, 0)
//	for _, v := range s {
//		// 如果这个字母之前出现过，直接将这个字母之前的组合一直合并到之前的
//		if index, ok := m[v]; ok {
//			res[index] += 1
//			// 将 res 合并 将 letter 合并，将 m 更换位置
//			for i := index + 1; i < len(res); i++ {
//				res[index] += res[i]
//				for _, v1 := range letter[i] {
//					m[v1] = index
//					letter[index] = append(letter[index], v1)
//				}
//			}
//			// 因为又添加了一个字符，res[index]也需要+1
//			res[index] += 1
//			res = res[:index+1]
//			letter = letter[:index+1]
//		} else {
//			res = append(res, 1)
//			// 记录这个字母的下表
//			m[v] = len(res) - 1
//			// 记录 这个下表的 字母串
//			letter = append(letter, []int32{v})
//		}
//	}
//	return res
//}
