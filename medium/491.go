package main

import "fmt"

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

var path []int
var res [][]int

func findSubsequences(nums []int) [][]int {
	path = []int{}
	res = [][]int{}

	backTracking(nums, 0)

	return res
}

func backTracking(nums []int, startIndex int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}
	// used数组要在这里建，为了树层去重
	//
	used := make([]int, 201)

	for i := startIndex; i < len(nums); i++ {

		if (len(path) > 0 && nums[i] < path[len(path)-1]) ||
			used[nums[i]+100] == 1 {
			continue
		}

		used[nums[i]+100] = 1

		path = append(path, nums[i])
		backTracking(nums, i+1)
		path = path[:len(path)-1]
	}
}
