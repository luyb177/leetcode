package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
//
// }
func removeElements(head *ListNode, val int) *ListNode {
	//虚拟头结点
	dummyHead := new(ListNode)
	dummyHead.Next = head

	current := dummyHead
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummyHead.Next
}
