package main

//func main() {
//
//}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var traversal func(node *TreeNode, count int) bool
	traversal = func(node *TreeNode, count int) bool {
		if node.Left == nil && node.Right == nil && count == 0 {
			return true
		}
		if node.Left == nil && node.Right == nil {
			return false
		}

		if node.Left != nil {
			count -= node.Left.Val
			if traversal(node.Left, count) {
				return true
			}
			count += node.Left.Val
		}
		if node.Right != nil {
			count -= node.Right.Val
			if traversal(node.Right, count) {
				return true
			}
			count += node.Right.Val
		}
		return false
	}

	return traversal(root, targetSum-root.Val)
}
