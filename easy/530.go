package main

import "math"

// func main() {
//
// }
//var pre *TreeNode
//var minSub int = math.MaxInt
//
//func getMinimumDifference(root *TreeNode) int {
//	// 中序遍历
//	if root == nil {
//		return minSub
//	}
//
//	// 左
//	temp := getMinimumDifference(root.Left)
//	if temp < minSub {
//		minSub = temp
//	}
//
//	// 中
//	if pre != nil {
//		temp = root.Val - pre.Val
//		if temp < minSub {
//			minSub = temp
//		}
//	}
//	pre = root
//
//	// 右
//	temp = getMinimumDifference(root.Right)
//	if temp < minSub {
//		minSub = temp
//	}
//	return minSub
//}

var pre *TreeNode
var minSub int

func traversal(root *TreeNode) {
	if root == nil {
		return
	}

	// 左
	traversal(root.Left)
	// 中
	if pre != nil {
		temp := root.Val - pre.Val
		if temp < minSub {
			minSub = temp
		}
	}
	pre = root
	// 右
	traversal(root.Right)
}

func getMinimumDifference(root *TreeNode) int {
	pre = nil
	minSub = math.MaxInt
	traversal(root)
	return minSub
}
