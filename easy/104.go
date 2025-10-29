package main

// func main() {
//
// }
//func maxDepth(root *TreeNode) int {
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
//			if node.Left != nil {
//				queue.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				queue.PushBack(node.Right)
//			}
//		}
//		res++
//	}
//	return res
//}
