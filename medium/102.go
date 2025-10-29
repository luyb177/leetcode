package main

import "container/list"

//func main() {
//
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历 -- 队列
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var tempRes []int     // 每一层的结果
		length := queue.Len() // 记录当前层的节点数
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) // 移出队列
			tempRes = append(tempRes, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, tempRes)
	}
	return res
}
