package main

//	func main() {
//		fmt.Println(totalFruit([]int{0, 1, 2, 2}))
//	}
func totalFruit(fruits []int) int {
	switch len(fruits) {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Max(fruits)
	}
}
func Max(fruits []int) int {
	i := 0
	j := 0
	temp_1 := fruits[i]
	temp_2 := -1
	result := len(fruits) + 1
	for ; j <= len(fruits)-1; j++ {
		//赋值
		for k := j; k <= len(fruits)-1; k++ {
			if fruits[k] != temp_1 {
				temp_2 = fruits[k]
				break
			}
		}
		if temp_2 == -1 {
			return len(fruits)
		}
		//第一次的值
		result = j - i + 1
		for k := j; k <= len(fruits)-1; k++ {
			if fruits[k] != temp_1 && fruits[k] != temp_2 { //遇到不同
				for temp := i + 1; temp < k; temp++ {
					if fruits[temp] == temp_1 {
						i = temp
					}
				}
			} else {
				if k := j - i + 1; k > result {
					result = k
				}
			}
		}
		temp_1, temp_2 = temp_2, -1
		i++

	}
	return result
}
