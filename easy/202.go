package main

//
//func main(){
//
//}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		var x = n % 10
		sum += x * x
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	m := make(map[int]struct{})
	sum := getSum(n)
	for {
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}
		sum = getSum(sum)
		if sum == 1 {
			return true
		}
	}

}

// 参考
func isHappy_1(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum_1(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
