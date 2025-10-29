package main

//	func main() {
//		fmt.Println(canConstruct("a", "b"))
//	}
func canConstruct(ransomNote string, magazine string) bool {
	//定义数组 哈希
	hash := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		num := magazine[i] - 'a'
		hash[num]++
	}
	//开始使用
	for i := 0; i < len(ransomNote); i++ {
		num := ransomNote[i] - 'a'
		hash[num]--
	}
	//检测结果
	for i := 0; i < len(hash); i++ {
		if hash[i] < 0 {
			return false
		}
	}
	return true
}
