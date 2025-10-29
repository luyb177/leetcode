package main

//
//func main() {
//
//}

func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 栈
	stack := make([]byte, 0)
	for i, _ := range s {
		if len(stack) == 0 {
			// 压栈
			stack = append(stack, s[i])
		} else {
			// peek
			peek := stack[len(stack)-1]
			if peek == s[i] {
				// 此时pop
				stack = stack[:len(stack)-1]
			} else {
				// 不想等的话压栈
				stack = append(stack, s[i])
			}
		}

	}
	return string(stack)
}
