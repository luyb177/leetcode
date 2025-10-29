package main

//func main() {
//
//}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}

	next := getNext(s)
	n := len(next)
	if next[n-1] != 0 {
		// 说明有最长相等前后缀
		if n%(n-next[n-1]) == 0 {
			// 说明有重复子串
			return true
		}
	}
	return false
}

// kmp
//func getNext(s string) []int{
//	if len(s) ==0 {
//		return nil
//	}
//
//	next := make([]int,len(s))
//	next[0] = 0
//	i := 0
//	for j := 1; j < len(s); j++{
//		for i >0 && s[i] != s[j]{
//			i = next[i - 1] // 持续回退
//		}
//		if s[i] == s[j]{
//			i++
//		}
//		next[j] = i
//	}
//	return next
//}
