package main

//func main() {
//	fmt.Println(reverseWords("the sky is blue"))
//}

func reverseWords(s string) string {

	result := []byte(s)
	n := len(result)

	// 使用快慢指针去除空格
	var slow = 0
	for fast := 0; fast < n; fast++ {
		if result[fast] != ' ' {
			if slow != 0 {
				// 不是第一个单词
				result[slow] = ' '
				slow++
			}

			// 读取直至下一个单词
			for fast < n && result[fast] != ' ' {
				result[slow] = result[fast]
				slow++
				fast++
			}
		}
	}

	// 截取新字符串 无空格
	result = result[:slow]
	n = len(result)
	// 使用双指针反转
	var left = 0
	var right = n - 1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	// 继续对单个单词进行反转
	slow = 0
	for fast := 0; fast < n; fast++ {

		if result[fast] == ' ' || fast == n-1 {
			// 这个是个界限，截取了一个单词
			left = slow
			right = fast - 1
			if fast == n-1 {
				right = fast
			}
			for left < right {
				result[left], result[right] = result[right], result[left]
				left++
				right--
			}
			// 更新一下数据
			slow = fast + 1 // 跳到下一个单词
		}
	}
	return string(result)

}
