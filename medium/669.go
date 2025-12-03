package main

// func main() {
//
// }
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		// 遍历右子树  返回符合区间的节点
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		// 遍历左子树 返回符合区间的节点
		return trimBST(root.Left, low, high)
	}
	//	接住符合条件的左右孩子
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
