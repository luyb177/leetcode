package main

import "strconv"

//func main() {
//
//}

func evalRPN(tokens []string) int {
	// 栈
	stack := make([]int, 0)

	for _, v := range tokens {
		vInt, err := strconv.Atoi(v)
		if err == nil {
			// 是数字 压栈
			stack = append(stack, vInt)
		} else {
			// 不是数字，是符号
			// 取出两个数 13 5  /
			num1 := stack[len(stack)-2] // 13
			num2 := stack[len(stack)-1] // 5
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}

	}
	return stack[0]
}
