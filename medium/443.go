package main

import (
	"strconv"
)

//func main() {
//	fmt.Println(compress([]byte("a")))
//}

func compress(chars []byte) int {
	slow := 0
	fast := 0
	n := len(chars)
	for fast < n {
		count := 0
		for fast < n && chars[fast] == chars[slow] {
			fast++
			count++
		}
		if count > 1 {
			if count >= 10 {
				// 变成字符串追加
				countStr := strconv.Itoa(count)
				for _, c := range countStr {
					slow++
					chars[slow] = byte(c)
				}
			} else {
				slow++
				chars[slow] = byte(count + '0')
			}
		}

		// 添加字符
		if fast < n {
			slow++
			chars[slow] = chars[fast]
		}
	}
	return slow + 1
}
