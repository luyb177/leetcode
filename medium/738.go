package main

import (
	"strconv"
)

//	func main() {
//		fmt.Println(monotoneIncreasingDigits(10))
//	}
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	// 标记后面都是 9
	flag := len(s)

	// 从低位向高位看是否是单调递减
	for i := len(s) - 2; i >= 0; i-- {
		// 有不符合的
		if s[i] > s[i+1] {
			s[i]--
			flag = i
		}
	}

	for i := flag + 1; i < len(s); i++ {
		s[i] = '9'
	}

	res, _ := strconv.Atoi(string(s))
	return res
}

// 贪心思路：
// 对于一个“单调递增”（即后面的数字 >= 前面的数字）的数字来说，
// 理想情况下每一位都尽可能保持原值，这样结果才最大。
//
// 从低位向高位遍历：
// 如果当前位 >= 前一位（高位 <= 低位），说明仍满足单调递增，直接保留。
// 一旦发现某一位破坏了递增关系：
//
//	例如 1221 中的最后一个 1，使得 2 > 1
//
// 就需要从前面开始调整：
//  1. 找到可以减 1 的位置
//  2. 将该位置减 1
//  3. 该位置后面的所有数字全部变成 9
//
// 为什么后面全变 9？
// 因为前面的数字已经确定变小了，
// 为了让整体结果尽可能大，后面的位应该全部取最大值 9。
// 这个实现方式上比较难以阅读
//func monotoneIncreasingDigits(n int) int {
//	if n == 0 {
//		return 0
//	}
//
//	// 例如：
//	// n = 1234 / 1221
//	// digits 中存储的是倒序数字
//	digits := make([]int, 0)
//	temp := n
//	for temp > 0 {
//		digits = append(digits, temp%10)
//		temp /= 10
//	}
//
//	res := make([]int, len(digits))
//	res[len(digits)-1] = digits[len(digits)-1]
//
//	for i := len(digits) - 2; i >= 0; i-- {
//
//		// 当前位仍满足单调递增，直接保留
//		if digits[i] >= res[i+1] {
//			res[i] = digits[i]
//		} else {
//
//			// 出现不满足递增的情况，需要向前调整
//			j := i + 1
//
//			// 找到一个可以减 1 的位置
//			for ; j <= len(digits)-1; j++ {
//
//				// 如果已经到最高位，直接减 1
//				if j == len(digits)-1 {
//					res[len(digits)-1]--
//					break
//				}
//
//				// 当前位减 1 后，仍满足单调递增
//				if res[j]-1 >= 0 && res[j]-1 >= res[j+1] {
//					res[j] = res[j] - 1
//					break
//				}
//			}
//
//			// 被调整位置之后的所有位全部变成 9
//			// 这样结果才能尽可能大
//			for j = j - 1; j >= 0; j-- {
//				res[j] = 9
//			}
//
//			break
//		}
//	}
//
//	// 将数字数组重新转回整数
//	resInt := res[len(res)-1]
//	for i := len(res) - 2; i >= 0; i-- {
//		resInt *= 10
//		resInt += res[i]
//	}
//
//	return resInt
//}
