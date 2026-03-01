package main

//	func main() {
//		letterCombinations("2")
//		fmt.Println(res)
//	}
//var res []string
//var path string
//
//var phoneMap = map[string]string{
//	"2": "abc",
//	"3": "def",
//	"4": "ghi",
//	"5": "jkl",
//	"6": "mno",
//	"7": "pqrs",
//	"8": "tuv",
//	"9": "wxyz",
//}
//
//func letterCombinations(digits string) []string {
//	res = []string{}
//	path = ""
//
//	if digits == "" {
//		return res
//	}
//
//	backtracking(digits, 0)
//	return res
//}
//
//func backtracking(digits string, index int) {
//	if index == len(digits) {
//		res = append(res, path)
//		return
//	}
//
//	letters := phoneMap[string(digits[index])]
//	for i := 0; i < len(letters); i++ {
//		path += string(letters[i])
//		backtracking(digits, index+1)
//		path = path[:len(path)-1]
//	}
//}

//var phoneMap = map[string]string{
//	"2": "abc",
//	"3": "def",
//	"4": "ghi",
//	"5": "jkl",
//	"6": "mno",
//	"7": "pqrs",
//	"8": "tuv",
//	"9": "wxyz",
//}

//var res []string
//var path string
//
//func letterCombinations(digits string) []string {
//	res = []string{}
//	path = ""
//
//	if digits == "" {
//		return res
//	}
//
//	backtracking(digits, 0)
//	return res
//}
//
//func backtracking(digits string, index int) {
//	if index == len(digits) {
//		res = append(res, path)
//		return
//	}
//
//	digit := digits[index]
//
//	words := phoneMap[string(digit)]
//
//	for _, word := range words {
//		path += string(word)
//		backtracking(digits, index+1)
//		path = path[:len(path)-1]
//	}
//}
