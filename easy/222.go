package main

import "container/list"

// func main() {
//
// }
func countNodes(root *TreeNode) int {
	// 层序遍历 - 之后写递归
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	var res = 1

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
				res++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
				res++
			}
		}
	}
	return res
}
