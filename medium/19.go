package main

//	func main() {
//		head := &ListNode{
//			Val:  1,
//			Next: nil,
//		}
//		fmt.Println(removeNthFromEnd(head, 1))
//	}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	//双指针
	slow := new(ListNode)
	fast := new(ListNode)
	dummy.Next = head
	slow = dummy
	fast = dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	} //让fast先行动几步
	//同步运动
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
