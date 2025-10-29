package main

//func main() {
//
//}

type ListNode struct {
	Val  int
	Next *ListNode //下一个节点的指针
}

func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head                             //虚拟头节点
	cur := dummy                                  //cur要在交换的两个节点前
	for cur.Next != nil && cur.Next.Next != nil { //条件也要有先后顺序
		//先保存节点
		temp := cur.Next
		temp1 := cur.Next.Next.Next

		//交换节点
		cur.Next = cur.Next.Next
		cur.Next.Next = temp
		temp.Next = temp1

		//移动cur
		cur = cur.Next.Next

	}
	return dummy.Next
}
