package main

import "container/list"

// func main() {
//
// }
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		var sum float64
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += float64(node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ave := sum / float64(length)
		res = append(res, ave)
	}
	return res
}
