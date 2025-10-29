package main

//	func main() {
//		s := "abc"
//		fmt.Println(s[0] - 'a')
//	}
func isAnagram(s string, t string) bool {
	//定义数组 哈希
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		num := s[i] - 'a'
		hash[num]++
	}
	for i := 0; i < len(t); i++ {
		num := t[i] - 'a'
		hash[num]--
	}
	for i := 0; i < len(hash); i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
