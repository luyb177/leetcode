package main

// func main() {
//
// }
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//记录size
	sizeA := Size(headA)
	sizeB := Size(headB)
	if sizeA > sizeB {
		return Move(sizeA-sizeB, headA, headB)
	} else {
		return Move(sizeB-sizeA, headB, headA)
	}
}
func Move(sub int, headA, headB *ListNode) *ListNode {
	//sizeA>sizeB
	curA := headA
	curB := headB
	result := new(ListNode)
	for i := 0; i < sub; i++ {
		//移动size大的
		curA = curA.Next
	}
	for ; curA != nil; curA = curA.Next {
		if curA == curB {
			result = curA
			return result
		}
		curB = curB.Next
	}
	return nil
}
func Size(head *ListNode) int {
	size := 0
	cur := head
	for ; cur != nil; cur = cur.Next {
		size++
	}
	return size
}
