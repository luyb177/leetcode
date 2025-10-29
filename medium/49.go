package main

import (
	"sort"
)

//func main() {
//	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
//}

// 数组哈希会超时
// 建议用map
// 直接删除str可能会出现越界问题
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	nums := make([]int, 0) // 存储已处理的索引

	for i := 0; i < len(strs); i++ {
		if contains(nums, i) { // 检查字符串是否已经分组过
			continue
		}

		record := make([]string, 0)
		record = append(record, strs[i])

		hash := make([]int, 26) // 每次初始化一个新的 hash 数组

		// 对第一个字符串进行 hash 处理
		reagain(strs[i], hash)

		for j := i + 1; j < len(strs); j++ {
			if contains(nums, j) { // 检查字符串是否已经分组过
				continue
			}

			// 重新生成 hash 来检查是否是异位词
			tempHash := make([]int, 26)
			reagain(strs[j], tempHash)

			// 如果两个字符串是字母异位词
			if equalHash(hash, tempHash) {
				record = append(record, strs[j])
				nums = append(nums, j)
			}
		}

		// 将当前分组的 record 添加到结果中，确保添加副本
		result = append(result, append([]string(nil), record...))
	}
	return result
}

// 辅助函数：检查一个索引是否已经在 nums 中
func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// 重置 hash 并计算字符串的字符频率
func reagain(str string, hash []int) {
	// 重置 hash
	for j := 0; j < len(hash); j++ {
		hash[j] = 0
	}
	// 设置 hash
	for j := 0; j < len(str); j++ {
		num := str[j] - 'a'
		hash[num]++
	}
}

// 比较两个 hash 是否相等
func equalHash(hash1, hash2 []int) bool {
	for i := 0; i < 26; i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}

// map
func groupAnagrams1(strs []string) [][]string {
	anagramMap := make(map[string][]string) // 用哈希表来存储字母异位词组

	for _, str := range strs {
		// 将字符串排序，作为哈希表的键
		sortedStr := sortString(str)

		// 将原始字符串添加到该键对应的值中
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// 将哈希表中的所有值添加到结果中
	result := make([][]string, 0)
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

// 排序字符串并返回结果
func sortString(s string) string {
	chars := []byte(s)
	sort.Sort(ByChar(chars))
	return string(chars)
}

// 排序的帮助函数
type ByChar []byte

func (b ByChar) Len() int           { return len(b) }
func (b ByChar) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByChar) Less(i, j int) bool { return b[i] < b[j] }
