package main

func main() {
	// fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}

//func lemonadeChange(bills []int) bool {
//	m := make(map[int]int)
//	m[5] = 0
//	m[10] = 0
//	m[20] = 0
//
//	for _, v := range bills {
//		switch v {
//		case 5:
//			m[5]++
//		case 10:
//			if m[5] > 0 {
//				m[5]--
//				m[10]++
//			} else {
//				return false
//			}
//		case 20:
//			if m[5] > 0 && m[10] > 0 {
//				m[5]--
//				m[10]--
//				m[20]++
//			} else if m[5] > 2 {
//				m[5] -= 3
//				m[20]++
//			} else {
//				return false
//			}
//		}
//	}
//	return true
//}
