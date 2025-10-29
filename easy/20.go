package main

//	func main() {
//		fmt.Println(isValid("([)]"))
//		//s := "()"
//		//s1 := "[]"
//		//s2 := "{}"
//		//
//		//fmt.Println(s[0], s[1])
//		//fmt.Println(s1[0], s1[1])
//		//fmt.Println(s2[0], s2[1])
//	}
// 暴力没做出来
//func isValid(s string) bool {
//	left := 0
//
//	for ; left < len(s)-1; left++ {
//		if s[left] != '(' && s[left] != '[' && s[left] != '{' {
//			continue
//		}
//		Flag := false
//		for right := left + 1; right < len(s); right++ {
//			flag := false
//			if s[right] != ')' && s[right] != ']' && s[right] != '}' {
//				continue
//			}
//			switch s[left] {
//			case '(':
//				if s[right] == ')' {
//					flag = true
//				}
//			case '[':
//				if s[right] == ']' {
//					flag = true
//				}
//			case '{':
//				if s[right] == '}' {
//					flag = true
//				}
//			}
//			if flag {
//				Flag = true
//				break
//			}
//		}
//		if !Flag {
//			return false
//		}
//	}
//	return true
//}

func isValid(s string) bool {
	// 提前判断
	if len(s)%2 != 0 {
		return false
	}

	// 提前记录
	m := make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	stack := make([]byte, 0)
	// 遍历
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			// 左括号直接入栈
			stack = append(stack, byte(v))
		default: // 右括号
			// 1 右边多余
			if len(stack) == 0 {
				return false
			}
			// 2 匹不匹配
			pop := stack[len(stack)-1]
			if pop != m[byte(v)] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	// 3 是否站内还有剩余
	return len(stack) == 0
}
