package main

// func main() {
//
// }
func sumOfLeftLeaves(root *TreeNode) int {
	// 后序遍历
	if root == nil {
		return 0
	}

	leftValue := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val // 该节点本身是左叶子节点
	}

	rightvalue := sumOfLeftLeaves(root.Right)

	return leftValue + rightvalue
}
