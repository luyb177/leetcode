package main

// func main() {
//
// }
//type Node struct {
//	Val      int
//	Children []*Node
//}

//func levelOrder1(root *Node) [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//
//	queue := list.New()
//	queue.PushBack(root)
//
//	for queue.Len() > 0 {
//		var tempRes []int
//		length := queue.Len()
//
//		for i := 0; i < length; i++ {
//			node := queue.Remove(queue.Front()).(*Node)
//			tempRes = append(tempRes, node.Val)
//			for _, c := range node.Children {
//				if c != nil {
//					queue.PushBack(c)
//				}
//			}
//		}
//		res = append(res, tempRes)
//	}
//	return res
//}
