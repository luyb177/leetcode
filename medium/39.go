package main

//func main() {
//
//}

var res [][]int
var path []int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	path = []int{}

	backtracking(candidates, target, 0)
	return res

}

func backtracking(candidates []int, target int, startIndex int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	} else if target < 0 {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtracking(candidates, target-candidates[i], i)
		path = path[:len(path)-1]
	}
}
