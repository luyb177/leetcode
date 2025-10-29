package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 创建队列
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		first := queue.Front()
		maxRes := first.Value.(*TreeNode).Val
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if maxRes < node.Val {
				maxRes = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, maxRes)
	}
	return res
}
