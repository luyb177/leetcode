package main

// func main() {
//
// }
// 二叉搜索树的有序性
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空节点就插入
	if root == nil {
		newNode := &TreeNode{Val: val}
		return newNode
	}

	// 遍历
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
