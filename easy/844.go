package main

import "fmt"

//	func main() {
//		fmt.Println(backspace("a##c"))
//		fmt.Println(backspace("#a#c"))
//
// }
func test() {
	a := "ab#c"
	var arr = []byte(a)
	var b string
	for i, v := range arr {
		if arr[i] != '#' {
			b += string(v)
			fmt.Println(b)
		}
	}
}
func backspace(s string) string {
	slow := 0
	fast := 0
	sArr := []byte(s)
	str := ""

	for ; fast <= len(sArr)-1; fast++ {
		if sArr[fast] == '#' {
			if slow > 0 {
				slow--
			}
		} else {
			sArr[slow] = sArr[fast]
			slow++
		}
	}
	for i := 0; i < slow; i++ {
		str += string(sArr[i])
	}
	return str
}
func backspaceCompare(s string, t string) bool {
	if backspace(s) == backspace(t) {
		return true
	} else {
		return false
	}
}
