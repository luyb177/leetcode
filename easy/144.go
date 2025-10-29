package main

//func main() {
//
//}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中 左 右
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中
		res = append(res, node.Val)
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return res
}
