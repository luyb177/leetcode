package main

//
//func main() {
//
//}

func reverseString(s []byte) {
	// 双指针
	var left = 0
	var right = len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}
