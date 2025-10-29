package main

import "fmt"

//	func main() {
//		buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
//		buildTree([]int{1, 2}, []int{2, 1})
//	}
func buildTree(inorder []int, postorder []int) *TreeNode {
	return traversal(inorder, postorder)
}

func traversal(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 后续遍历最后一个节点
	rootVal := postorder[len(postorder)-1]
	root := TreeNode{
		Val: rootVal,
	}

	if len(postorder) == 1 {
		return &root
	}

	// 切割一下 中序数组
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	leftInorder := inorder[:i]
	fmt.Printf("leftInorder: %v\n", leftInorder)
	rightInorder := inorder[i+1:]
	fmt.Printf("rightInorder: %v\n", rightInorder)

	// 依据中序 切割后序
	leftPostorder := postorder[:len(leftInorder)]
	fmt.Printf("leftPostorder: %v\n", leftPostorder)
	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]
	fmt.Printf("rightPostorder: %v\n", rightPostorder)

	root.Left = traversal(leftInorder, leftPostorder)
	root.Right = traversal(rightInorder, rightPostorder)
	return &root
}
