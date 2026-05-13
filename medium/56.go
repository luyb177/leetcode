package main

import (
	"sort"
)

//
//func main() {
//	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
//}

// Note：
// 这里有两个要点就是，
// 1. 目前代码实现的是在下一次循环的时候才会append上一个区间，所以最后的时候需要再append一下
// 2. 在合并的时候，要判断当前的left和当前的区间的起点哪个更小，确保留下来的区间是一个较长的区间
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	left := intervals[len(intervals)-1][0]
	right := intervals[len(intervals)-1][1]
	for i := len(intervals) - 2; i >= 0; i-- {
		if left <= intervals[i][1] {
			// 合并
			if left > intervals[i][0] {
				left = intervals[i][0]
			}
		} else {
			res = append(res, []int{left, right})
			right = intervals[i][1]
			left = intervals[i][0]
		}
	}
	res = append(res, []int{left, right})
	return res
}
