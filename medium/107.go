package main

import "container/list"

//func main() {
//
//}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	// 队列
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var tempRes []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
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

	// 将结果翻转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
