package main

import "strconv"

// func main() {
//
// }
func binaryTreePaths(root *TreeNode) []string {
	var res = []string{}
	var path = []int{}
	if root == nil {
		return res
	}
	traversal(root, &path, &res)
	return res
}

//
//// 前序遍历
//func traversal(cur *TreeNode, path *[]int, result *[]string) {
//	// 中
//	*path = append(*path, cur.Val)
//
//	// 判断是否是叶子节点
//	if cur.Left == nil && cur.Right == nil {
//		// 是叶子节点
//		sPath := ""
//		for i := 0; i < len(*path)-1; i++ {
//			sPath += strconv.Itoa((*path)[i])
//			sPath += "."
//		}
//		sPath += strconv.Itoa((*path)[len(*path)-1])
//		*result = append(*result, sPath)
//	}
//
//	// 非叶子节点
//	// 左
//	if cur.Left != nil {
//		traversal(cur.Left, path, result)
//		*path = (*path)[:len(*path)-1] // 回溯
//	}
//
//	if cur.Right != nil {
//		traversal(cur.Right, path, result)
//		// 回溯
//		*path = (*path)[:len(*path)-1] // 回溯
//	}
//}
