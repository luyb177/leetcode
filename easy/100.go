package main

import "container/list"

// func main() {
//
// }
// FS/DFS 每次只取一对，而不是四个一组。
// 检查两棵树是否相同，只需要成对比较节点
// p.Left 和 q.Left
// p.Right 和 q.Right
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := list.New()
	queue.PushBack(p)
	queue.PushBack(q)

	for queue.Len() > 0 {
		node1 := queue.Remove(queue.Front()).(*TreeNode)
		node2 := queue.Remove(queue.Front()).(*TreeNode)

		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		queue.PushBack(node1.Left)
		queue.PushBack(node2.Left)
		queue.PushBack(node1.Right)
		queue.PushBack(node2.Right)
	}
	return true
}
