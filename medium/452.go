package main

import (
	"sort"
)

//func main() {
//	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
//}

// 按一个纬度
// 从左到右
func findMinArrowShots(points [][]int) int {
	// 按照气球的左边界从小到大
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	leftCount := 1
	curXstart := points[0][1]
	for i := 1; i < len(points); i++ {
		if curXstart >= points[i][0] {
			if curXstart > points[i][1] {
				curXstart = points[i][1]
			}
		} else {
			curXstart = points[i][1]
			leftCount++
		}
	}

	rightCount := 1
	curXstart = points[len(points)-1][0]
	for i := len(points) - 2; i >= 0; i-- {
		if curXstart > points[i][1] {
			curXstart = points[i][0]
			rightCount++
		}
	}
	if leftCount > rightCount {
		return rightCount
	} else {
		return leftCount
	}
}
