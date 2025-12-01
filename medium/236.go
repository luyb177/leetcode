package main

// func main() {
//
// }
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == q || root == p || root == nil {
		return root
	}

	// 左
	left := lowestCommonAncestor(root.Left, p, q)

	// 右
	right := lowestCommonAncestor(root.Right, p, q)

	// 中
	if right != nil && left != nil {
		return root
	}

	if right != nil && left == nil {
		return right
	} else if right == nil && left != nil {
		return left
	} else {
		return nil
	}

}
