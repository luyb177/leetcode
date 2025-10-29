package main

import "container/list"

//func main() {
//
//}

// 层序遍历 左右指针均为空的时候到达了最低点
//func minDepth(root *TreeNode) int {
//	var res = 0
//	if root == nil {
//		return res
//	}
//
//	queue := list.New()
//	queue.PushBack(root)
//
//	for queue.Len() > 0 {
//		length := queue.Len()
//		for i := 0; i < length; i++ {
//			node := queue.Remove(queue.Front()).(*TreeNode)
//			if node.Left == nil && node.Right == nil {
//				res++
//				return res
//			}
//			if node.Left != nil {
//				queue.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushBack(node.Left)
//			}
//
//		}
//		res++
//	}
//	return res
//}

func minDepth(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return res + 1
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res++
	}

	return res
}
