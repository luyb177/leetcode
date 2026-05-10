package main

//func main() {
//	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
//}

func partitionLabels(s string) []int {
	res := make([]int, 0)
	// 记录 这个字母上次出现的位置
	m := make(map[int32]int)
	// 记录 这个位置的 字母
	letter := make([][]int32, 0)
	for _, v := range s {
		// 如果这个字母之前出现过，直接将这个字母之前的组合一直合并到之前的
		if index, ok := m[v]; ok {
			res[index] += 1
			// 将 res 合并 将 letter 合并，将 m 更换位置
			for i := index + 1; i < len(res); i++ {
				res[index] += res[i]
				for _, v1 := range letter[i] {
					m[v1] = index
					letter[index] = append(letter[index], v1)
				}
			}
			// 因为又添加了一个字符，res[index]也需要+1
			res[index] += 1
			res = res[:index+1]
			letter = letter[:index+1]
		} else {
			res = append(res, 1)
			// 记录这个字母的下表
			m[v] = len(res) - 1
			// 记录 这个下表的 字母串
			letter = append(letter, []int32{v})
		}
	}
	return res
}
