package main

import "container/list"

//func main() {
//
//}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == length-1 {
				res = append(res, node.Val) // 将当前层的最后一个节点加入结果
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
