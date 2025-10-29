package main

import "container/list"

// func main() {
//
// }
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		tempArr := make([]*Node, 0, length)
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr = append(tempArr, node)
		}
		for i := 0; i < length-1; i++ {
			tempArr[i].Next = tempArr[i+1]
		}
	}
	return root
}
