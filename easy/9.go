package main

import (
	"strconv"
)

//
//func main() {
//	fmt.Println(isPalindrome(121))
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	nums := []byte(strconv.Itoa(x))
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] != nums[right] {
			return false
		}
		left++
		right--
	}
	return true
}
