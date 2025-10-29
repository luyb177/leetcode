package main

import "container/list"

// func main() {
//
// }
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 队列
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		// 判断

		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil {
			continue
		} else if left.Val != right.Val {
			return false
		}

		// 左右节点不为空且值相等
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
