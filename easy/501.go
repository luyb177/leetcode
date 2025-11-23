package main

//func main() {

//}

// var pre *TreeNode
var count int
var res []int
var maxCount int

func findMode(root *TreeNode) []int {
	count = 0
	maxCount = 0
	pre = nil
	res = []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左
		traversal(node.Left)

		// 中
		if pre == nil {
			count = 1
		} else if pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		pre = node

		if count > maxCount {
			res = []int{node.Val}
			maxCount = count
		} else if count == maxCount {
			res = append(res, node.Val)
		}

		traversal(node.Right)
	}
	traversal(root)
	return res
}
