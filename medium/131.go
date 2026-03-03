package main

//func main() {
//	fmt.Printf("%v", partition("aab"))
//}
//
//var res [][]string
//var path []string
//
//func partition(s string) [][]string {
//	res = [][]string{}
//	path = []string{}
//
//	backTracking(s, 0)
//	return res
//}
//
//func backTracking(s string, startIndex int) {
//	if startIndex >= len(s) {
//		temp := make([]string, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//
//	for i := startIndex; i < len(s); i++ {
//		// 其中 s[startIndex:i]就是要截取的子串
//		if isPalindrome(s, startIndex, i) {
//			str := s[startIndex : i+1]
//			path = append(path, str)
//		} else {
//			continue
//		}
//		backTracking(s, i+1)
//		path = path[:len(path)-1]
//	}
//}
//
//func isPalindrome(s string, startIndex int, endIndex int) bool {
//	for i, j := startIndex, endIndex; i <= j; i, j = i+1, j-1 {
//		if s[i] != s[j] {
//			return false
//		}
//	}
//	return true
//}
