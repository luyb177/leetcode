package main

//	func main() {
//		fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
//	}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	left := 0
	right := 0
	count := 200
	for ; right < len(strs); right++ {
		//循环最小的那个
		temp := 0
		m := min(len(strs[left]), len(strs[right]))
		for j := 0; j < m; j++ {
			if strs[left][j] == strs[right][j] {
				temp++
			} else {
				break
			}
		}
		if count > temp {
			count = temp
		}
	}
	return strs[0][:count]
}
