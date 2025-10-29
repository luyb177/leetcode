package main

import "container/list"

// func main() {
//
// }
// 使用前序和后序遍历都是可以的，这里选择使用层序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}
