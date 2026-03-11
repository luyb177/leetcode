package main

//
////func main() {}
//var path string
//var res []string
//
//func restoreIpAddresses(s string) []string {
//	path = ""
//	res = []string{}
//
//	backTracking(s,0)
//	return res
//}
//
//
//func backTracking (s string, startIndex int){
//	if len(path) == len(s) + 3 {
//		// 不能只看字符串长度来区分分了几段，需要再有一个参数
//		temp := path
//		res = append(res,temp)
//		return
//	}
//
//	// 截取的字符串的长度是
//	// [startIndex:i+1] 其中i < startIndex + 3 && len(s) - i >= 4 - len(path) - i
//	for i := startIndex; i < startIndex + 3 && len(s) - i >= 4 - len(path) + i ; i++ {
//		str := s[startIndex:i+1]
//		strInt, _ := strconv.Atoi(str)
//		if strInt > 255 {
//			continue
//		}
//		path = path + "." + str
//		backTracking(s,i+1)
//		path = path[:len(s)-len(str) - 1]
//	}
//}
//
//// path 最后使用strings.Join来结合，使用切片也方便回溯
//var path []string
//var res []string
//
//func restoreIpAddresses(s string) []string {
//	path = []string{}
//	res = []string{}
//
//	backTracking(s, 0, 1)
//	return res
//}
//
//// segment 已经切的段数
//func backTracking(s string, startIndex int, segment int) {
//	if segment == 4 {
//		// 切成 4 段，且字符用完
//		if startIndex == len(s) {
//			res = append(res, strings.Join(path, "."))
//		}
//		return
//	}
//
//	// 计算本次切分的最小值,确保本次截取不超过3，且剩余字符够用
//	remainSegment := 4 - segment
//	num1 := len(s) - startIndex - remainSegment
//	num2 := startIndex + 3
//	num := Min(num1, num2)
//
//	// todo 这里的 i 的范围其实不太对
//	for i := startIndex; i < num; i++ {
//		// 截取一下
//		str := s[startIndex : i+1]
//
//		// 不包含前导0
//		// todo 这里会把 0 也给排除
//		if str[0] == '0' {
//			break
//		}
//
//		// 不能大于255
//		strInt, _ := strconv.Atoi(str)
//		if strInt > 255 {
//			break
//		}
//
//		path = append(path, str)
//		backTracking(s, i+1, segment+1)
//		path = path[:len(path)-1]
//	}
//
//}
//
//func Min(a int, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//
//var path []string
//var res []string
//
//func restoreIpAddresses(s string) []string {
//	path = []string{}
//	res = []string{}
//
//	backTracking(s, 0, 0)
//	return res
//}
//
//func backTracking(s string, startIndex int, segment int) {
//	// 分成4段，且字符用完了
//	if segment == 4 {
//		// 这里是多进入一次递归，让segment == 4 的时候，startIndex == len(s) 了
//		if startIndex == len(s) {
//			res = append(res, strings.Join(path, "."))
//		}
//		return
//	}
//
//	// 让剩余的字符数满足剩余的段数，每段至少1个字符，最多3个字符
//	remainSegment := 4 - segment
//	remainChars := len(s) - startIndex
//	if remainChars < remainSegment || remainChars > remainSegment*3 {
//		return
//	}
//
//	// 同理 截取字符串s[startIndex:i+1]
//	for i := startIndex; i < startIndex+3 && i < len(s); i++ {
//		str := s[startIndex : i+1]
//
//		// 不能有前导0，0除外
//		if len(str) > 1 && str[0] == '0' {
//			break
//		}
//
//		// 不能大于255
//		num, _ := strconv.Atoi(str)
//		if num > 255 {
//			break
//		}
//
//		path = append(path, str)
//		backTracking(s, i+1, segment+1)
//		path = path[:len(path)-1]
//	}
//}
