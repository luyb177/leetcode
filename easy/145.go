package main

// func main() {
//
// }
// 左 右 中
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
