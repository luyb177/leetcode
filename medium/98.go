package main

// func main() {
//
// }
var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return validate(root)
}

func validate(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !validate(root.Left) {
		return false
	}

	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root

	return validate(root.Right)
}
