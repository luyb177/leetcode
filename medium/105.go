package main

// func main() {
//
// }
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := TreeNode{
		Val: rootVal,
	}

	if len(preorder) == 1 {
		return &root
	}

	// 切割
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}

	leftInorder := inorder[:i]
	rightInorder := inorder[i+1:]

	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return &root
}
