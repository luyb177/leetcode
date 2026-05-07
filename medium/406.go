package main

import "sort"

// 先处理高个
// 先处理不会被后面影响的人
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		// 身高相等，先处理前面k小的
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 身高不相等，先处理高个的
		return people[i][0] > people[j][0]
	})

	var res [][]int

	for _, person := range people {
		// 插入到k的位置
		k := person[1]
		res = append(res[:k], append([][]int{person}, res[k:]...)...)
	}
	return res
}

// 要先确定一个纬度，然后在这个纬度上进行排序，最后再插入到结果中
// 已经将高个的排序了，那么再将矮个的插入并不影响 k
// 这时候矮个的k已经就是在插入的位置
