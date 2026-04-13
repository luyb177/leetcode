package main

//func main() {
//	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
//}

// 超出时间限制
//func canCompleteCircuit(gas []int, cost []int) int {
//	temp := make([]int,len(gas))
//	for i := range gas {
//		temp[i] = gas[i] - cost[i]
//	}
//
//	sum := 0
//	for i := 0; i < len(temp); i++ {
//		sum += temp[i]
//	}
//	if sum < 0 {
//		return -1
//	}
//
//	for i := 0; i < len(temp); i++ {
//		if temp[i] < 0 {
//			continue
//		}
//		sum := 0
//		for j := 0; j < len(temp); j++ {
//			num := (i+j) % len(temp)
//			sum += temp[num]
//			if sum < 0 {
//				break
//			}
//		}
//		if sum >= 0 {
//			return i
//		}
//	}
//	return -1
//}

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0   // 当前区间
	totalSum := 0 // 总区间
	index := 0
	for i := range gas {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			index = i + 1
		}
	}
	if totalSum < 0 {
		return -1
	}
	return index
}
