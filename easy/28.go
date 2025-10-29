package main

//func main() {
//	fmt.Println(strStr("mississippi", "issip"))
//}
//func strStr(haystack string, needle string) int {
//	needleByte := []byte(needle)
//	if len(needleByte) == 0 {
//		return 0
//	}
//	haystackByte := []byte(haystack)
//	result := -1
//	next := getNext(needle)
//
//	j := 0
//	for i := 0; i < len(haystack); i++ {
//
//
//
//		if haystackByte[i] == needleByte[j] {
//			j++
//			if j == len(needleByte) {
//				return i + 1 - j
//			}
//		} else {
//			if j > 0 {
//				j = next[j-1]
//				i-- // 循环结束保持i不变
//			}
//		}
//	}
//	return result
//}

//	func getNext(s string) []int {
//		sByte := []byte(s)
//		j := 0
//		next := make([]int, len(sByte))
//		next[0] = j
//
//		for i := 1; i < len(sByte); i++ {
//			for j > 0 && sByte[i] != sByte[j] {
//				j = next[j-1]
//			}
//			if sByte[i] == sByte[j] {
//				j++
//			}
//			next[i] = j
//		}
//		return next
//	}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := getNext(needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func getNext(needle string) []int {
	result := []byte(needle)
	if len(result) == 0 {
		return nil
	}
	next := make([]int, len(needle))

	// 初始化
	var i = 0
	next[0] = 0

	for j := 1; j < len(result); j++ {
		// 不相等的话 i一直后退
		for i > 0 && result[j] != result[i] {
			i = next[i-1]
		}

		// 相等的话
		if result[j] == result[i] {
			i++
		}
		next[j] = i
	}
	return next
}
