package main

import "sort"

// Note: 会发现按终点从小到大排序的分支会少一点
// 是其已经将区间的大小排序了，前面的起点比后面的起点校，前面的终点比后面的终点校
// 只需要比较前面的终点和后面的起点就行了
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}

		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	curXstart := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if curXstart > intervals[i][0] {
			res++
		} else {
			curXstart = intervals[i][1]
		}
	}
	return res
}

// Note: 按起点排序
// 这一版是也比较了 curStart 和 intervals[i][1] 的大小，确保留下来的是一个较小的区间
//func eraseOverlapIntervals(intervals [][]int) int {
//	sort.Slice(intervals, func(i, j int) bool {
//		if intervals[i][0] == intervals[j][0] {
//			return intervals[i][1] < intervals[j][1]
//		}
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	res := 0
//	curStart := intervals[0][1]
//	for i := 1; i < len(intervals); i++ {
//		if curStart > intervals[i][0] {
//			if curStart >= intervals[i][1] {
//				curStart = intervals[i][1]
//			}
//			res++
//		} else {
//			curStart = intervals[i][1]
//		}
//	}
//	return res
//}

// Note: 这里是按照起点来排序的，那么就会出现一种情况
// [1,100] [2,3] [3,4] 此时正确结果应该删掉[1,100]，但是此时代码逻辑是跳过下一个
// 修改方案就是每次保留区间是保留终点最小的区间
//func eraseOverlapIntervals(intervals [][]int) int {
//	sort.Slice(intervals,func(i, j int) bool {
//		if intervals[i][0] == intervals[j][0] {
//			return intervals[i][1] < intervals[j][1]
//		}
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	res := 0
//	curStart := intervals[0][1]
//	for i := 1; i < len(intervals); i++ {
//		if curStart > intervals[i][0] {
//			res++
//		} else {
//			curStart = intervals[i][1]
//		}
//	}
//	return res
//}
