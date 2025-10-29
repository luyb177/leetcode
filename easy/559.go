package main

import (
	"container/list"
)

// func main() {
//
// }
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	var res = 0

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for _, child := range node.Children {
				if child != nil {
					queue.PushBack(child)
				}
			}
		}
		res++
	}
	return res
}
