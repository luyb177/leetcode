package main

// func main() {
//
// }
func isBalanced(root *TreeNode) bool {
	if getHeight(root) != -1 {
		return true
	}
	return false
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if Abs(leftHeight-rightHeight) > 1 {
		return -1
	} else {
		return 1 + Max(leftHeight, rightHeight)
	}
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
