package main

func isPerfectSquare(num int) bool {
	flag := true
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		} else {
			flag = false
		}
	}
	return flag
}
