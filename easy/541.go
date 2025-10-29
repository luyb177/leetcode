package main

//func main() {
//	fmt.Println(reverseStr("abcdefg", 8))
//}

func reverseStr(s string, k int) string {
	result := []byte(s)
	for i := 0; i < len(result)-1; i += 2 * k {
		// 截取数组
		var left = i
		var right = i + k - 1

		// 当截取的数组越界时
		if right > len(result)-1 {
			right = len(result) - 1
		}

		for left < right {
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}
	}
	return string(result)
}
